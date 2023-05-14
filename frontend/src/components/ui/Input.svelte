<script>
	export let value = '';
	export let label = '';
	export let name = '';
	export let pending = false;
	export let errors = {};
	export let onInput = () => {};

	$: messages = errors[name] || [];
</script>

<label class:pending class={$$props.class}>
	<div class="row">
		<strong class="col-xs-4">{label}</strong>
		{#if messages.length}
			<span class="col-xs-8 error-container">{messages[0]}</span>
		{/if}
	</div>
	<input type="text" {name} bind:value {...$$props} on:input={onInput} />
</label>
