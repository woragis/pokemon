<script lang="ts">
	import { register } from '$lib/api/auth';
	import { goto } from '$app/navigation';

	let first_name = '';
	let last_name = '';
	let username = '';
	let email = '';
	let password = '';
	let error: string | null = null;

	async function handleRegister() {
		error = null;
		try {
			await register({ username, email, password, first_name, last_name });
			goto('/login'); // redirect after register
		} catch (e: any) {
			error = e.message;
		}
	}
</script>

<form on:submit|preventDefault={handleRegister}>
	<input type="text" bind:value={first_name} placeholder="First name" required />
	<input type="text" bind:value={last_name} placeholder="Last name" required />
	<input type="text" bind:value={username} placeholder="Username" required />
	<input type="email" bind:value={email} placeholder="Email" required />
	<input type="password" bind:value={password} placeholder="Password" required />
	<button type="submit">Register</button>
</form>

{#if error}
	<p style="color: red;">{error}</p>
{/if}
