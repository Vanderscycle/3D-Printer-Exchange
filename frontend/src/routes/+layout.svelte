<script>
	import '../app.css';
	import Modal from 'svelte-simple-modal';
	import { uiState } from '$stores/ui';
	let { darkMode } = uiState;
	import Navbar from '$components/ui/Navbar.svelte';
	import Footer from '$components/ui/Footer.svelte';
</script>

<div class={$darkMode === true ? 'dark h-full' : 'h-full'}>
	<Navbar />
	<main class="dark:bg-darkBg bg-Bg h-screen dark:text-darkText text-Text">
		<Modal>
			<slot />
		</Modal>
	</main>
	<Footer />
</div>

<style>
	@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono&display=swap');
</style>
