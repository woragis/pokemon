<script lang="ts">
  import { login } from '@/api/auth';
  import { goto } from '$app/navigation';

  let email = '';
  let password = '';
  let error: string | null = null;

  async function handleLogin() {
    error = null;
    try {
      const data = await login({ email, password });
      // save token, e.g. localStorage or a store
      localStorage.setItem('token', data.token);
      goto('/dashboard'); // redirect after login
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
