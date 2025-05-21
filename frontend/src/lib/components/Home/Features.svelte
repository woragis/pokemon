<script lang='ts'>
    export let featuredGuides
    export let pokemonOfTheDay
    export let regions
</script>

<!-- Featured Guides -->
<section class="py-12 max-w-5xl mx-auto px-4">
<h2 class="text-3xl font-bold mb-6 text-center">⭐ Featured Guides</h2>
<ul class="grid gap-6 md:grid-cols-3">
    {#each featuredGuides as guide}
    <li class="bg-white p-4 rounded shadow hover:bg-yellow-50 transition">
        <a href={guide.link} class="text-xl font-semibold text-red-600">{guide.title}</a>
    </li>
    {/each}
</ul>
</section>

<!-- Pokémon of the Day + Regions -->
<section class="py-12 bg-white">
<div class="max-w-6xl mx-auto grid md:grid-cols-2 gap-8 px-4">
    <!-- Pokémon of the Day -->
    <div class="text-center">
    <h2 class="text-2xl font-bold mb-4">📅 Pokémon of the Day</h2>
    <img src={pokemonOfTheDay.sprite} alt={pokemonOfTheDay.name} class="mx-auto w-32 h-32" />
    <h3 class="text-xl font-semibold mt-2">{pokemonOfTheDay.name}</h3>
    <p class="text-sm mt-1">{pokemonOfTheDay.fact}</p>
    </div>

    <!-- Regions -->
    <div>
    <h2 class="text-2xl font-bold mb-4 text-center">🗺️ Browse by Region</h2>
    <div class="flex flex-wrap justify-center gap-4">
        {#each regions as region}
        <a href={"/region/" + region.toLowerCase()} class="bg-yellow-200 text-gray-800 font-medium px-4 py-2 rounded hover:bg-yellow-300">
            {region}
        </a>
        {/each}
    </div>
    </div>
</div>
</section>