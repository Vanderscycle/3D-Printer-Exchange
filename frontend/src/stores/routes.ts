import type { NavigationButton } from '$lib/interfaces';
import type { Writable } from 'svelte/store';
import { writable } from 'svelte/store';

export const routes: Writable<NavigationButton[]> = writable([
	{ url: '/', name: '/Home' },
	{ url: '/browse', name: '/Browse' },
	{ url: '/about', name: '/About' }
]);
