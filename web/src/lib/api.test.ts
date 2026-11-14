import { create, getOption } from '@bufbuild/protobuf';
import { Code, ConnectError, createRouterTransport } from '@connectrpc/connect';
import { expect, test, vi } from 'vitest';
import {
	field as fieldRules,
	FieldPathElementSchema,
	FieldPathSchema,
	ViolationSchema,
	ViolationsSchema
} from './gen/buf/validate/validate_pb';
import { ErrorCode, ErrorDetailSchema } from './gen/casa/v1/error_pb';
import { HouseholdService } from './gen/casa/v1/household_pb';
import { IdentityService } from './gen/casa/v1/identity_pb';
import { ClientKind, RequestContextSchema, type RequestContext } from './gen/casa/v1/request_pb';
import { CLIENT_VERSION, createApi, ctx, errorCode, failureOf, fieldViolations } from './api';

const uuid = '0b6f3c2e-8d4a-4f1b-9c7e-2a5d6e8f1b3c';

/** An API whose server records the request context of every call. */
function recordingApi() {
	const seen: RequestContext[] = [];
	const transport = createRouterTransport(({ service }) => {
		service(IdentityService, {
			signIn(request) {
				if (request.requestContext) seen.push(request.requestContext);
				return {};
			}
		});
		service(HouseholdService, {
			listMemberships(request) {
				if (request.requestContext) seen.push(request.requestContext);
				return {};
			}
		});
	});
	return { api: createApi(transport), seen };
}

function failed(
	code: Code,
	...details: NonNullable<ConstructorParameters<typeof ConnectError>[3]>
) {
	return new ConnectError('failed', code, undefined, details);
}

function detail(code: ErrorCode) {
	return { desc: ErrorDetailSchema, value: { code } };
}

test('SEC-169 (d): every request carries client_kind web and the build version', async () => {
	const { api, seen } = recordingApi();
	await api.household.listMemberships({ requestContext: ctx(true) });
	await api.identity.signIn({
		requestContext: { ...ctx(true), clientVersion: '9.9.9' },
		email: 'a@example.org',
		password: 'x'
	});
	expect(seen).toHaveLength(2);
	for (const context of seen) {
		expect(context.clientKind).toBe(ClientKind.WEB);
		expect(context.clientVersion).toBe(CLIENT_VERSION);
	}
});

test('SEC-169 (d): the build version has the form the contract requires of client_version', () => {
	const rules = getOption(RequestContextSchema.field.clientVersion, fieldRules);
	expect(rules.type.case).toBe('string');
	const pattern = rules.type.case === 'string' ? rules.type.value.pattern : '';
	expect(pattern).not.toBe('');
	expect(CLIENT_VERSION).toMatch(new RegExp(pattern));
});

test('SEC-165 (g): user_initiated and the idempotency key reach the server unchanged', async () => {
	const { api, seen } = recordingApi();
	await api.household.listMemberships({ requestContext: ctx(false) });
	await api.household.listMemberships({ requestContext: ctx(true, uuid) });
	expect(seen.map((c) => [c.userInitiated, c.idempotencyKey])).toEqual([
		[false, ''],
		[true, uuid]
	]);
});

test('ADR-013 item 4: the domain code comes from ErrorDetail, not from the Connect code', () => {
	expect(errorCode(failed(Code.FailedPrecondition, detail(ErrorCode.NOTICE_UNREAD)))).toBe(
		ErrorCode.NOTICE_UNREAD
	);
	expect(errorCode(failed(Code.FailedPrecondition))).toBeUndefined();
	expect(errorCode(new TypeError('Failed to fetch'))).toBeUndefined();
});

test('ADR-030 item 10: failures are sorted into offline, server, session and rejected', () => {
	const fetchFailed = ConnectError.from(new TypeError('Failed to fetch'));
	expect(failureOf(fetchFailed)).toBe('offline');
	expect(failureOf(failed(Code.DeadlineExceeded))).toBe('offline');
	expect(failureOf(failed(Code.Unavailable))).toBe('server');
	expect(failureOf(failed(Code.Internal))).toBe('server');
	expect(failureOf(failed(Code.Unknown))).toBe('server');
	expect(failureOf(failed(Code.Unavailable, detail(ErrorCode.WRITE_FROZEN)))).toBe('server');
	expect(failureOf(failed(Code.Aborted, detail(ErrorCode.RETRY_EXHAUSTED)))).toBe('server');
	expect(failureOf(failed(Code.Unauthenticated))).toBe('session');
	expect(failureOf(failed(Code.Unauthenticated, detail(ErrorCode.SESSION_EXPIRED)))).toBe(
		'session'
	);
	expect(failureOf(failed(Code.Unauthenticated, detail(ErrorCode.SESSION_REVOKED)))).toBe(
		'session'
	);
	expect(failureOf(failed(Code.Unauthenticated, detail(ErrorCode.CREDENTIALS_REJECTED)))).toBe(
		'rejected'
	);
	expect(failureOf(failed(Code.ResourceExhausted, detail(ErrorCode.RATE_LIMITED)))).toBe(
		'rejected'
	);
	expect(failureOf(failed(Code.InvalidArgument))).toBe('rejected');
});

test('ADR-030 item 10: any failure while the browser is offline is offline', () => {
	vi.stubGlobal('navigator', { onLine: false });
	expect(failureOf(failed(Code.Unavailable))).toBe('offline');
});

test('ADR-013 item 4: field violations map each field path to its rule id', () => {
	const violation = (ruleId: string, ...names: string[]) =>
		create(ViolationSchema, {
			ruleId,
			field: create(FieldPathSchema, {
				elements: names.map((fieldName) => create(FieldPathElementSchema, { fieldName }))
			})
		});
	const error = failed(Code.InvalidArgument, {
		desc: ViolationsSchema,
		value: {
			violations: [
				violation('name.visible_length', 'name'),
				violation('string.pattern', 'request_context', 'client_version')
			]
		}
	});
	expect(fieldViolations(error)).toEqual(
		new Map([
			['name', 'name.visible_length'],
			['request_context.client_version', 'string.pattern']
		])
	);
	expect(fieldViolations(failed(Code.InvalidArgument)).size).toBe(0);
});
