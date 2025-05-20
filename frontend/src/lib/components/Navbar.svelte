<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  // Mock user logged-in state
  let loggedIn = false;
  let username = '';

  onMount(() => {
    // Here you could check auth status from cookie/localStorage or API
    loggedIn = Boolean(localStorage.getItem('token'));
    username = loggedIn ? 'AshKetchum' : '';
  });

  function logout() {
    localStorage.removeItem('token');
    loggedIn = false;
    username = '';
    goto('/login');
  }
</script>

<nav class="navbar">
  <ul class="nav-list">
    <li><a href="/">Home</a></li>
    <li><a href="/blog">Blog</a></li>
    <li><a href="/pokedex">Pokedex</a></li>
    <li><a href="/walkthroughs">Walkthroughs</a></li>
    <!-- <li><a href="/regions">Regions</a></li> -->
    {#if loggedIn}
      <li><a href="/profile">{username}</a></li>
      <li><button on:click={logout}>Logout</button></li>
    {:else}
      <li><a href="/auth/login">Login</a></li>
      <li><a href="/auth/register">Register</a></li>
    {/if}
  </ul>
</nav>

<style>
  .navbar {
    background-color: #2a75bb;
    padding: 1rem 2rem;
  }
  .nav-list {
    list-style: none;
    display: flex;
    gap: 1.5rem;
    margin: 0;
    padding: 0;
    align-items: center;
  }
  a {
    color: white;
    text-decoration: none;
    font-weight: 600;
  }
  a:hover {
    text-decoration: underline;
  }
  button {
    background: none;
    border: none;
    color: white;
    font-weight: 600;
    cursor: pointer;
  }
  button:hover {
    text-decoration: underline;
  }
</style>
