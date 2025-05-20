<script lang="ts">
	import '@/app.css';
	import Navbar from '@/components/Navbar.svelte';
	import Footer from '@/components/Footer.svelte';

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
