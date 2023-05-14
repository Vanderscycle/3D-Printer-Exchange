<script>
	// import classNames from "vest/classNames";
 // import validate from './validate';

	import Input from '$components/ui/Input.svelte';
	import Button from '$components/ui/Button.svelte';
	import Checkbox from '$components/ui/Checkbox.svelte';

	export let values = {};
	export let submit = () => {};

	let errors = {};
	let warnings = {};
	// we copy the supplied for values to our form state
	// this is a good practice and helps us with resetting the form
	let formState = { ...values };

	let usernamePending = false;

	// initial empty validation state
 // let result = validate.get();
 //
 // 	const check = ({ target: { name } }) => runValidate(name);
 //
 // 	const runValidate = (name) => {
 // 		result = validate(formState, name);
 //
 // 		// extract errors and warnings
 // 		errors = result.getErrors();
 // 		warnings = result.getWarnings();
 //
 // 	};
 //
 // const reset = () => {
 // 		validate.reset();
 // 		result = validate.get();
 // 		errors = {};
 // 		warnings = {};
 // 		formState = { ...values };
 // 	};
 //
 // 	const onSubmit = () => {
 // 		runValidate();
 // 		if (result.hasErrors()) return;
 // 		submit(formState);
 // 		// reset form state
 // 		reset();
 // 	};
 //
	// Vest css class helper utility
	// $: cn = classNames(result, {
	//   warning: "warning",
	//   invalid: "invalid",
	//   valid: "valid"
	// });
	//
	// reactive variable for the submit button
 // $: disabled = result.hasErrors();
</script>

<form id="example-form" class="col-xs-10 col-lg-6">
	<h2>Authentication</h2>
	<Input
		name="username"
		label="Username"
		pending={usernamePending}
		placeholder="try: ealush or codechips"
		{errors}
		bind:value={formState.username}
	/>
	<Input
		name="password"
		label="Password"
		errors={{ ...errors, ...warnings }}
		bind:value={formState.password}
	/>
	<Input
		name="confirm_password"
		label="Confirm Password"
		{errors}
		bind:value={formState.confirm_password}
	/>
	<Input name="email" label="Email" {errors} bind:value={formState.email}/>
	<Checkbox
		name="tos"
		label="I have read and agreed to the terms of service."
		bind:checked={formState.tos}
	/>
	<footer>
		<Button  >Submit</Button>
	</footer>
</form>

	<!-- onInput={check} -->
