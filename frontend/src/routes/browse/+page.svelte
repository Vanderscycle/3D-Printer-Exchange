<script lang="ts">
	import QRCode from '$components/ui/QRcode.svelte';
	import Button from '$components/ui/Button.svelte';
	import { fade } from 'svelte/transition';

	let message: string = 'hello';
	let newSellNotice: Boolean = false;
	import AuthForm from '$components/forms/Sellnotice.svelte';

	// const submit = (values) => console.log(values);

	const values = {
		username: '',
		password: '',
		confirm_password: '',
		email: ''
	};
</script>

<div class="">
	Title

	<Button callbackFn={() => (newSellNotice = !newSellNotice)}>Button</Button>
	{#if newSellNotice}
		<div transition:fade><AuthForm {values} submit={() => console.log(values)} /></div>
	{/if}

	<QRCode data={message} />
</div>
