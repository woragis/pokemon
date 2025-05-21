<script lang="ts">
	import { login } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import { setCookie } from '$lib/api';

	let email = '';
	let password = '';
	let error: string | null = null;

	async function handleLogin() {
		error = null;
		try {
			const data = await login({ email, password });
			goto('/dashboard');
		} catch (e: any) {
			error = e.message;
		}
	}
</script>

<form on:submit|preventDefault={handleLogin}>
	<input type="email" bind:value={email} placeholder="Email" required />
	<input type="password" bind:value={password} placeholder="Password" required />
	<button type="submit">Login</button>
</form>

{#if error}
	<p style="color: red;">{error}</p>
{/if}
