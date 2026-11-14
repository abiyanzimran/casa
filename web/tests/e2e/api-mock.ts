// Answers the app's Connect calls in Playwright with the generated contract (plan "Uji alur"):
// requests are read and answers written with the generated schemas, in the binary format the
// app uses (plan decision 2), and errors as Connect JSON errors with binary details.
import {
	create,
	fromBinary,
	toBinary,
	type DescMessage,
	type DescMethodUnary,
	type MessageInitShape,
	type MessageShape
} from '@bufbuild/protobuf';
import { Code } from '@connectrpc/connect';
import { expect, type Page } from '@playwright/test';
import { ViolationsSchema } from '../../src/lib/gen/buf/validate/validate_pb';
import { ErrorDetailSchema, type ErrorCode } from '../../src/lib/gen/casa/v1/error_pb';

/** A failed answer; build it with rpcError(). */
export interface RpcError {
	code: Code;
	body: string;
}

/** One request the app sent to a mocked method. */
export type RpcCall<I extends DescMessage> = MessageShape<I>;

// Connect protocol: HTTP status per code (connectrpc.com/docs/protocol, "Error Codes").
const httpStatus: Record<Code, number> = {
	[Code.Canceled]: 499,
	[Code.Unknown]: 500,
	[Code.InvalidArgument]: 400,
	[Code.DeadlineExceeded]: 504,
	[Code.NotFound]: 404,
	[Code.AlreadyExists]: 409,
	[Code.PermissionDenied]: 403,
	[Code.ResourceExhausted]: 429,
	[Code.FailedPrecondition]: 400,
	[Code.Aborted]: 409,
	[Code.OutOfRange]: 400,
	[Code.Unimplemented]: 501,
	[Code.Internal]: 500,
	[Code.Unavailable]: 503,
	[Code.DataLoss]: 500,
	[Code.Unauthenticated]: 401
};

function detail(type: string, value: Uint8Array): { type: string; value: string } {
	// Connect sends details as unpadded standard base64.
	return { type, value: Buffer.from(value).toString('base64').replace(/=+$/, '') };
}

/** A Connect error with the domain code and the buf.validate violations the server sends. */
export function rpcError(
	code: Code,
	errorCode?: ErrorCode,
	violations?: { field: string; ruleId: string }[]
): RpcError {
	const details = [];
	if (errorCode !== undefined) {
		const value = toBinary(ErrorDetailSchema, create(ErrorDetailSchema, { code: errorCode }));
		details.push(detail(ErrorDetailSchema.typeName, value));
	}
	if (violations) {
		const message = create(ViolationsSchema, {
			violations: violations.map(({ field, ruleId }) => ({
				field: { elements: field.split('.').map((fieldName) => ({ fieldName })) },
				ruleId
			}))
		});
		details.push(detail(ViolationsSchema.typeName, toBinary(ViolationsSchema, message)));
	}
	// The wire name of a code is its enum name in snake case (DeadlineExceeded → deadline_exceeded).
	const name = Code[code].replace(/(?<!^)([A-Z])/g, '_$1').toLowerCase();
	return { code, body: JSON.stringify({ code: name, message: name, details }) };
}

function isRpcError(reply: object): reply is RpcError {
	return 'code' in reply && 'body' in reply && typeof reply.body === 'string';
}

/**
 * Answers every call of method with reply(request) and returns the requests as they arrive.
 * A reply may wait (a slow server); a request the page has given up on is then dropped.
 */
export async function mockRpc<I extends DescMessage, O extends DescMessage>(
	page: Page,
	method: DescMethodUnary<I, O>,
	reply: (
		request: MessageShape<I>
	) => MessageInitShape<O> | RpcError | Promise<MessageInitShape<O> | RpcError>
): Promise<RpcCall<I>[]> {
	const calls: RpcCall<I>[] = [];
	await page.route(`**/${method.parent.typeName}/${method.name}`, async (route) => {
		const request = fromBinary(method.input, route.request().postDataBuffer() ?? new Uint8Array());
		calls.push(request);
		const answer = await reply(request);
		const response = isRpcError(answer)
			? { status: httpStatus[answer.code], contentType: 'application/json', body: answer.body }
			: {
					status: 200,
					contentType: 'application/proto',
					body: Buffer.from(toBinary(method.output, create(method.output, answer)))
				};
		await route.fulfill(response).catch(() => {
			// The page aborted the request (its deadline passed) while the reply waited.
		});
	});
	return calls;
}

/** Makes any Connect call without a mock fail the test; register before the mocks. */
export async function failUnmocked(page: Page): Promise<void> {
	await page.route(/\/casa\.v1\.\w+\/\w+$/, async (route) => {
		expect.soft(new URL(route.request().url()).pathname, 'a Connect call without a mock').toBe('');
		await route.abort();
	});
}
