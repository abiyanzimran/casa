// The one API module of the web client (ADR-030 items 2, 9): clients generated from the
// contract, the client version on every request, and the readers of a failed call.
import type { MessageInitShape } from '@bufbuild/protobuf';
import { Code, ConnectError, createClient, type Client, type Transport } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { ViolationsSchema } from './gen/buf/validate/validate_pb';
import { ErrorCode, ErrorDetailSchema } from './gen/casa/v1/error_pb';
import { HouseholdService } from './gen/casa/v1/household_pb';
import { IdentityService } from './gen/casa/v1/identity_pb';
import { ClientKind, RequestContextSchema } from './gen/casa/v1/request_pb';

/** The build version, `version` in web/package.json (ADR-030 item 9; SEC-169 (d)). */
export const CLIENT_VERSION: string = __CASA_VERSION__;

type RequestContextInit = MessageInitShape<typeof RequestContextSchema>;

/**
 * Stamps client_kind web and the build version on every request, over whatever the caller
 * set, and leaves user_initiated and the idempotency key as the caller gave them
 * (SEC-169 (d), SEC-165 (g)). Every request has request_context as field 1 (14 §4.6).
 */
function clientIdentity(transport: Transport): Transport {
	return {
		unary: (method, signal, timeoutMs, header, input, contextValues) => {
			const request = input as { requestContext?: RequestContextInit };
			const stamped = {
				...input,
				requestContext: {
					...request.requestContext,
					clientKind: ClientKind.WEB,
					clientVersion: CLIENT_VERSION
				}
			};
			return transport.unary(method, signal, timeoutMs, header, stamped, contextValues);
		},
		// ponytail: the contract has no streaming RPC; stamp here too when one is added.
		stream: transport.stream.bind(transport)
	};
}

/** The generated clients over `transport`, every request stamped by clientIdentity. */
export function createApi(transport: Transport): {
	identity: Client<typeof IdentityService>;
	household: Client<typeof HouseholdService>;
} {
	const stamped = clientIdentity(transport);
	return {
		identity: createClient(IdentityService, stamped),
		household: createClient(HouseholdService, stamped)
	};
}

/**
 * Same origin, binary Protobuf, the browser's own fetch and cookies: Origin and
 * Sec-Fetch-Site come from the browser, and no script reads the session cookie
 * (SEC-161 (b), SEC-162 (b); plan decision 2). The root-relative base resolves to
 * location.origin in the browser and lets this module load where there is no location
 * (Vitest).
 */
function webTransport(): Transport {
	return createConnectTransport({ baseUrl: '/', useBinaryFormat: true });
}

export const api = createApi(webTransport());

/**
 * The request context a caller sets: whether a person's action caused the call (SEC-165 (g))
 * and, for a write command, the key of its intent (ADR-005). The kind and version are
 * stamped by clientIdentity.
 */
export function ctx(userInitiated: boolean, idempotencyKey?: string): RequestContextInit {
	return { userInitiated, idempotencyKey };
}

function connectError(error: unknown): ConnectError | undefined {
	return error instanceof ConnectError ? error : undefined;
}

/** The domain code of a failed call, from ErrorDetail, never from the Connect code (ADR-013 item 4). */
export function errorCode(error: unknown): ErrorCode | undefined {
	return connectError(error)?.findDetails(ErrorDetailSchema)[0]?.code;
}

/** How a screen treats a failed call (ADR-030 item 10). */
export type Failure = 'offline' | 'server' | 'session' | 'rejected';

const serverCodes = new Set([Code.Unavailable, Code.Internal, Code.Unknown]);

function unreachable(error: unknown): boolean {
	// Only an explicit false is offline: Node's navigator (Vitest) has no onLine at all.
	// eslint-disable-next-line @typescript-eslint/no-unnecessary-boolean-literal-compare -- see above
	if (navigator.onLine === false) return true;
	// connect-web wraps a fetch that never reached the server as Unknown with the TypeError
	// as its cause; a deadline set with timeoutMs ends as DeadlineExceeded.
	if (error instanceof TypeError) return true;
	const failed = connectError(error);
	return failed?.cause instanceof TypeError || failed?.code === Code.DeadlineExceeded;
}

export function failureOf(error: unknown): Failure {
	if (unreachable(error)) return 'offline';
	const code = errorCode(error);
	switch (code) {
		case ErrorCode.WRITE_FROZEN:
		case ErrorCode.RETRY_EXHAUSTED:
			return 'server';
		case ErrorCode.SESSION_EXPIRED:
		case ErrorCode.SESSION_REVOKED:
			return 'session';
		case undefined: {
			const connectCode = ConnectError.from(error).code;
			if (serverCodes.has(connectCode)) return 'server';
			if (connectCode === Code.Unauthenticated) return 'session';
			return 'rejected';
		}
		default:
			return 'rejected';
	}
}

/** Each field path of a buf.validate violation, mapped to its rule id (ADR-013 item 4). */
export function fieldViolations(error: unknown): Map<string, string> {
	const violations = new Map<string, string>();
	for (const detail of connectError(error)?.findDetails(ViolationsSchema) ?? []) {
		for (const violation of detail.violations) {
			const path = (violation.field?.elements ?? []).map((e) => e.fieldName).join('.');
			if (!violations.has(path)) violations.set(path, violation.ruleId);
		}
	}
	return violations;
}
