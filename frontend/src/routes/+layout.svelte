<script lang="ts">
	import '../app.css';
	import '$lib/styles/app.css';
	import Footer from '$lib/components/Footer.svelte';
	import Navbar from '$lib/components/Navbar.svelte';

	let { children } = $props();
</script>

<Navbar />
<main id="main-content">
{@render children()}
</main>
<Footer />

<style>
	#main-content {
		min-height: calc(100vh - 80px);
	}
</style>
