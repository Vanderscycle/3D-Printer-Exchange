<script lang="ts">
	import Heroicon from '$components/icons/heroicons/Heroicon.svelte';
	import { sun as outlineSun } from '$components/icons/heroicons/outline';
	import { moon as outlineMoon } from '$components/icons/heroicons/outline';
	import { uiState } from '$stores/ui';
	import { routes } from '$stores/routes';
	import Button from './Button.svelte';
	import { goto } from '$app/navigation';
	const unwantedRoutes: string[] = ['/philosophy', '/testbench'];

	let toggleState: boolean = false;
	let { darkMode } = uiState;
	let filteredNav = $routes.filter((i) => !unwantedRoutes.includes(i.url));

	$: uiState.darkMode.set(!toggleState);
</script>

<nav class="dark:bg-darkGui bg-Gui">
	<div class="inline-flex w-full">
		<div class="mt-3 mx-2">
			{#if toggleState}
				<Heroicon icon={outlineSun} class={'text-Yellow'} />
			{:else}
				<Heroicon icon={outlineMoon} />
			{/if}
		</div>
		{#each filteredNav as navBtn}
			<Button callbackFn={() => goto(navBtn.url)}>{navBtn.name}</Button>{/each}
	</div>
</nav>
