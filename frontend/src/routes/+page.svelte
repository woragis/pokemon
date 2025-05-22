<script>
	import HeroSection from '$lib/components/home/HeroSection.svelte';
	import FeatureSection from '$lib/components/home/FeatureSection.svelte';
	import PokemonShowcase from '$lib/components/home/PokemonShowcase.svelte';
	import StatsSection from '$lib/components/home/StatsSection.svelte';
</script>

<div class="min-h-screen bg-gradient-to-br from-yellow-100 to-red-100 font-sans text-gray-900">
	<HeroSection />
	<FeatureSection />
	<PokemonShowcase />
	<StatsSection />
</div>
