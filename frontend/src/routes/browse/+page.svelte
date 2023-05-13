<script lang="ts">
	import QRCode from '$components/ui/QRCode.svelte';
	let message: string = 'hello';
</script>

<div class="">
	Title

	<QRCode data={message} />
</div>
