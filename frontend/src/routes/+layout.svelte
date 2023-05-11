<script>
	import '../app.css';
	import Modal from 'svelte-simple-modal';
	import { uiState } from '$stores/ui';
	let { darkMode } = uiState;
	import Navbar from '$components/ui/Navbar.svelte';
</script>

<div class={$darkMode === true ? 'dark h-full' : 'h-full'}>
	<nav class="dark:bg-darkGui bg-Gui">
		<Navbar />
	</nav>
	<main class="dark:bg-darkBg bg-Bg h-screen">
		<Modal>
			<slot />
		</Modal>
	</main>
</div>

<style>
	@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono&display=swap');
</style>
