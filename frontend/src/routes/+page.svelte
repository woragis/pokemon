<script>
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
  <!-- Hero -->
  <section class="text-center py-16 bg-red-500 text-white">
    <h1 class="text-5xl font-extrabold mb-4">🔥 Catch 'Em All</h1>
    <p class="text-xl mb-6">Explore the world of Pokémon with our detailed guides and tools!</p>
    <div class="space-x-4">
      <a href="/guides" class="bg-white text-red-600 font-semibold py-2 px-6 rounded shadow hover:bg-red-100">
        Explore Guides
      </a>
      <a href="/search" class="bg-red-600 border border-white py-2 px-6 rounded shadow hover:bg-white hover:text-red-600">
        Search Pokémon
      </a>
    </div>
  </section>

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

  <!-- News / Updates -->
  <section class="py-12 max-w-5xl mx-auto px-4">
    <h2 class="text-2xl font-bold mb-4 text-center">📰 Latest Updates</h2>
    <ul class="list-disc list-inside text-gray-700">
      <li>New guide: <a href="#" class="text-red-600 underline">Ash vs Cynthia Battle</a></li>
      <li>Sinnoh region maps added!</li>
    </ul>
  </section>

  <!-- Call to Action -->
  <section class="py-12 bg-red-100 text-center">
    <h2 class="text-2xl font-bold mb-4">Join the Trainer Community!</h2>
    <p class="mb-6">Sign up to save your favorite Pokémon, write reviews, and more.</p>
    <a href="/register" class="bg-red-600 text-white px-6 py-2 rounded hover:bg-red-700">
      Sign Up
    </a>
  </section>

  <!-- Footer -->
  <footer class="py-6 bg-gray-800 text-white text-center text-sm">
    <p>© 2025 Pokémon Guide. Not affiliated with Nintendo or Game Freak.</p>
    <div class="mt-2 space-x-4">
      <a href="/about" class="underline">About</a>
      <a href="/terms" class="underline">Terms</a>
      <a href="https://github.com/yourrepo" class="underline">GitHub</a>
    </div>
  </footer>
</div>
