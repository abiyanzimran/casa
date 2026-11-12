import { defineConfig } from 'eslint/config';
import svelte from 'eslint-plugin-svelte';
import ts from 'typescript-eslint';

// SEC-167: no HTML sinks and no string-to-code (ADR-030 item 7). Lint is the
// main guard, because Trusted Types is enforced only from Safari 26.
const sec167 = 'SEC-167: HTML sinks are not allowed; render text through Svelte markup.';

/* eslint-disable no-restricted-syntax -- this is the list of banned names itself */
const htmlSinkProperties = [
	'innerHTML',
	'outerHTML',
	'insertAdjacentHTML',
	'setHTMLUnsafe',
	'parseHTMLUnsafe',
	'createContextualFragment',
	'srcdoc',
	'parseFromString'
];
/* eslint-enable no-restricted-syntax */
const htmlSinkPattern = `/^(${htmlSinkProperties.join('|')})$/`;

export default defineConfig(
	{ ignores: ['build/', '.svelte-kit/', 'playwright-report/', 'test-results/'] },
	ts.configs.strictTypeChecked,
	svelte.configs.recommended,
	{
		languageOptions: {
			parserOptions: {
				projectService: true,
				tsconfigRootDir: import.meta.dirname,
				extraFileExtensions: ['.svelte']
			}
		}
	},
	{
		files: ['**/*.svelte', '**/*.svelte.ts'],
		languageOptions: { parserOptions: { parser: ts.parser } }
	},
	{
		rules: {
			'no-eval': 'error',
			'no-new-func': 'error',
			// typescript-eslint replaces the core no-implied-eval with this
			// type-aware rule (string passed to setTimeout, setInterval, Function).
			'@typescript-eslint/no-implied-eval': 'error',
			'svelte/no-at-html-tags': 'error',
			// Covers both el.innerHTML and el['innerHTML'] (static computed keys).
			'no-restricted-properties': [
				'error',
				...htmlSinkProperties.map((property) => ({ property, message: sec167 })),
				// window.eval(...): no browser globals are declared, so no-eval alone misses it.
				{ property: 'eval', message: sec167 },
				{ object: 'document', property: 'write', message: sec167 },
				{ object: 'document', property: 'writeln', message: sec167 }
			],
			'no-restricted-syntax': [
				'error',
				{
					// window.document.write(...) and the like.
					selector:
						"MemberExpression[object.property.name='document'][property.name=/^(write|writeln)$/]",
					message: sec167
				},
				{
					// <div bind:innerHTML={...}> and the other HTML properties.
					selector: `SvelteDirective[kind='Binding'][key.name.name=${htmlSinkPattern}]`,
					message: sec167
				},
				{
					// Object.assign(el, { innerHTML: s }) and object keys in general.
					selector: `Property[key.name=${htmlSinkPattern}]`,
					message: sec167
				},
				{
					// Reflect.set(el, 'innerHTML', s), el.setAttribute('srcdoc', s),
					// const key = 'innerHTML'.
					selector: `Literal[value=${htmlSinkPattern}]`,
					message: sec167
				},
				{
					// <iframe srcdoc="...">
					selector: "SvelteAttribute[key.name='srcdoc']",
					message: sec167
				}
			]
		}
	}
);
