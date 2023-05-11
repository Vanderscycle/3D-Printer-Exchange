<script>
	import '../app.css';
	import Modal from 'svelte-simple-modal';
 // import { uiState } from '$stores/ui';
 // 	let { darkMode } = uiState;
</script>

<main class="dark:bg-darkBg bg-Bg h-screen">
	<Modal>
		<slot />
	</Modal>
</main>

<style>
	@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono&display=swap');
</style>
