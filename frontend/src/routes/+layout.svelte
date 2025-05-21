<script lang="ts">
	import '../app.css';
	import '$lib/styles/app.css';
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';

	let { children } = $props();
</script>

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
