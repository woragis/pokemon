<script>
	import Cta from '$lib/components/Home/CTA.svelte';
  import Features from '$lib/components/Home/Features.svelte';
  import Hero from '$lib/components/Home/Hero.svelte';
	import News from '$lib/components/Home/News.svelte';

  const featuredGuides = [
    { title: "How to Beat Elite Four", link: "#" },
    { title: "Best Starters by Region", link: "#" },
    { title: "All Gym Leaders Ranked", link: "#" }
  ];

  const regions = ["Kanto", "Johto", "Hoenn", "Sinnoh", "Unova", "Kalos"];

  const pokemonOfTheDay = {
    name: "Pikachu",
    sprite: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png",
    fact: "Pikachu stores electricity in its cheeks and releases it in lightning-based attacks."
  };
</script>

<div class="min-h-screen bg-gradient-to-br from-yellow-100 to-red-100 text-gray-900 font-sans">
   <Hero />
   <Features featuredGuides={featuredGuides} pokemonOfTheDay={pokemonOfTheDay} regions={regions} />
   <News />
   <Cta />
</div>
