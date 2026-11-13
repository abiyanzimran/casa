// Red proof only: a second place that registers the service worker and makes a policy.
export async function registerElsewhere(): Promise<void> {
	await navigator.serviceWorker.register('/service-worker.js', { type: 'module', scope: '/' });
}
