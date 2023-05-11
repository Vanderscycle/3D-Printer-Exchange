<script lang="ts">
	import Heroicon from '$components/icons/heroicons/Heroicon.svelte';
	import { sun as outlineSun } from '$components/icons/heroicons/outline';
	import { moon as outlineMoon } from '$components/icons/heroicons/outline';
	let toggleState: boolean = false;
	import { uiState } from '$stores/ui';
	let { darkMode } = uiState;
</script>

<template class="">
	NAVBARRR!!
	<div class="inline-flex w-full">
		<div class="mt-3 mx-2">
			{#if toggleState}
				<Heroicon icon={outlineSun} class={'text-Yellow'} />
			{:else}
				<Heroicon icon={outlineMoon} />
			{/if}
		</div>
	</div></template
>
