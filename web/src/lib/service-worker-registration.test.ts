import { beforeEach, expect, test, vi } from 'vitest';

interface PolicyRules {
	createScriptURL: (input: string) => string;
}

const url = '/service-worker.js';
let register = vi.fn();

beforeEach(() => {
	vi.resetModules();
	register = vi.fn(() => Promise.resolve({}));
	vi.stubGlobal('navigator', { serviceWorker: { register } });
});

function stubTrustedTypes() {
	const policies: { name: string; rules: PolicyRules }[] = [];
	const createPolicy = vi.fn((name: string, rules: PolicyRules) => {
		policies.push({ name, rules });
		return { createScriptURL: (input: string) => ({ trusted: rules.createScriptURL(input) }) };
	});
	vi.stubGlobal('trustedTypes', { createPolicy });
	return { createPolicy, policies };
}

test('SEC-167 (c): one casa-sw policy with only createScriptURL is made when the module loads', async () => {
	const { createPolicy, policies } = stubTrustedTypes();
	await import('./service-worker-registration');
	expect(createPolicy).toHaveBeenCalledTimes(1);
	expect(policies[0].name).toBe('casa-sw');
	expect(Object.keys(policies[0].rules)).toEqual(['createScriptURL']);
});

test('SEC-167 (c): the casa-sw policy passes the service worker URL and throws for any other', async () => {
	const { policies } = stubTrustedTypes();
	await import('./service-worker-registration');
	const rules = policies[0].rules;
	expect(rules.createScriptURL(url)).toBe(url);
	for (const other of [
		'/service-worker.js?probe=1',
		'/service-worker.js#x',
		'https://example.com/service-worker.js',
		'//example.com/service-worker.js',
		'/other.js',
		' /service-worker.js'
	]) {
		expect(() => rules.createScriptURL(other)).toThrow();
	}
});

test('SEC-167 (c): the service worker is registered through the policy with fixed options', async () => {
	stubTrustedTypes();
	const { registerServiceWorker } = await import('./service-worker-registration');
	await registerServiceWorker();
	expect(register).toHaveBeenCalledWith({ trusted: url }, { type: 'module', scope: '/' });
});

test('SEC-167 (c): without Trusted Types the same constant URL is registered as a string', async () => {
	vi.stubGlobal('trustedTypes', undefined);
	const { registerServiceWorker } = await import('./service-worker-registration');
	await registerServiceWorker();
	expect(register).toHaveBeenCalledWith(url, { type: 'module', scope: '/' });
});
