<script lang="ts">
	import '../app.css';
	import '$lib/styles/app.css';
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';
	import { onMount } from 'svelte';
	import { getUser } from '$lib/store/user';

	let { children } = $props();
	onMount(() => {
		getUser();
	});
</script>

<svelte:head>
	<title>Pokemon</title>
	<meta name="description" content="This is the homepage of my Svelte app." />
	<meta property="og:title" content="Pokemon" />
	<meta property="og:description" content="This is the homepage of my Svelte app." />
</svelte:head>

<div class="flex min-h-screen flex-col">
	<Navbar />
	<main id="main-content" class="flex-grow pt-16">
		{@render children()}
	</main>
	<Footer />
</div>

<style>
	#main-content {
		min-height: calc(100vh - 80px);
	}
</style>
