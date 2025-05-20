<script>
	let pokemonList = [
		{ id: 1, name: 'Bulbasaur' },
		{ id: 4, name: 'Charmander' },
		{ id: 7, name: 'Squirtle' }
	];
</script>

<h1 class="text-2xl font-bold mb-4">Pokédex</h1>

<ul class="grid grid-cols-2 md:grid-cols-4 gap-4">
	{#each pokemonList as p}
		<li class="bg-white rounded-lg shadow p-4 hover:bg-gray-50">
			<a href={`/pokedex/${p.id}`}>
				<h2 class="font-semibold">{p.name}</h2>
				<p>ID: {p.id}</p>
			</a>
		</li>
	{/each}
</ul>
