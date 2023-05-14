<script>
	export let name = '';
	export let checked = false;
	export let label = '';
	export let onChange = () => {};
</script>

<label class={`v-center ${$$props.class}`}>
	<input type="checkbox" {name} bind:checked on:change={onChange} />
	<small>{label}</small>
</label>
