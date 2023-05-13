<script lang="ts">
	// let testdata = {
	// 	id: 1,
	// 	name: "User",
	// 	email: "user@gmail.com"
	// };
	// console.log(JSON.stringify(testdata));
	//lib from https://github.com/soldair/node-qrcode
	import QRCode from 'qrcode';
	import { onMount } from 'svelte';

	let canvas: HTMLCanvasElement;
	export let data: string = '';
	export let simpleModalMsg: string = 'Follow link';
	export let hyperlink: boolean = false;
	onMount(() => {
		QRCode.toCanvas(canvas, data, { errorCorrectionLevel: 'H' }, (error) => {
			if (error) {
				console.error(error);
			} else {
				console.log('success!');
			}
		});
	});
</script>

<div class="text-center flex flex-col justify-center items-center dark:text-darkText text-Text">
	<canvas
		class="rounded-xl cursor-pointer"
		bind:this={canvas}
		on:click={() => console.log('svelte allows you to do plenty of cool stuff, eh!')}
	/>
</div>

<style lang="postcss"></style>
