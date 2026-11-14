// What a form shows for a failed call (08 §7.6 rules for the whole table, §13.2): a sentence on
// a field for a buf.validate violation the form names, otherwise one sentence above the button.
import { tick } from 'svelte';
import { errorCode, failureOf, fieldViolations } from './api';
import { copy } from './copy';
import { ErrorCode } from './gen/casa/v1/error_pb';

/**
 * The sentence above the button for a failure that is no one field's (08 §7.6, §13.2): too many
 * tries, no connection, or the server's fault, which is logged as unexpected (ES §4: the error,
 * never the form's contents).
 */
export function actionFailure(error: unknown): string {
	if (errorCode(error) === ErrorCode.RATE_LIMITED) return copy.common.rateLimited;
	if (failureOf(error) === 'offline') return copy.serverAction.offline;
	console.error(error);
	return copy.serverAction.server;
}

/**
 * The field errors of a buf.validate answer, by field, from the sentences the form gives per
 * field and rule (ADR-013 item 4: the rule is the server's, the client only words it).
 * Undefined when the answer names no field and rule the form knows.
 */
export function fieldErrors(
	error: unknown,
	sentences: Partial<Record<string, Partial<Record<string, string>>>>
): Partial<Record<string, string>> | undefined {
	const found: Partial<Record<string, string>> = {};
	for (const [field, rule] of fieldViolations(error)) {
		const sentence = sentences[field]?.[rule];
		if (sentence) found[field] = sentence;
	}
	return Object.keys(found).length > 0 ? found : undefined;
}

/**
 * After a failed submit with field errors, the first field marked invalid takes focus once the
 * errors render, so its error is read (08 §25.5 "Formulir"; WCAG 4.1.3).
 */
export async function focusFirstInvalid(form: HTMLFormElement | undefined): Promise<void> {
	await tick();
	form?.querySelector<HTMLElement>('[aria-invalid="true"]')?.focus();
}
