// Red proof for SEC-167: every form below must be reported by web-lint.
export function sinks(el: HTMLElement, s: string): void {
	el.innerHTML = s;
	el['innerHTML'] = s;
	el.outerHTML = s;
	el.insertAdjacentHTML('beforeend', s);
	document.write(s);
	document.writeln(s);
	window.document.write(s);
	setTimeout('alert(1)', 1);
	new Function(s);
	eval(s);
}
