<script lang="ts">
	interface Pokemon {
		id: number;
		name: string;
		image: string;
		types: string[];
	}

	const typeColors: Record<string, string> = {
		Normal: 'bg-gray-300',
		Fire: 'bg-red-500',
		Water: 'bg-blue-500',
		Electric: 'bg-yellow-400',
		Grass: 'bg-green-500',
		Ice: 'bg-blue-200',
		Fighting: 'bg-red-700',
		Poison: 'bg-purple-500',
		Ground: 'bg-yellow-600',
		Flying: 'bg-indigo-300',
		Psychic: 'bg-pink-500',
		Bug: 'bg-green-400',
		Rock: 'bg-yellow-700',
		Ghost: 'bg-purple-700',
		Dragon: 'bg-indigo-600',
		Dark: 'bg-gray-700',
		Steel: 'bg-gray-400',
		Fairy: 'bg-pink-300'
	};

	const pokemonList: Pokemon[] = [
		{
			id: 1,
			name: 'Bulbasaur',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/1.png',
			types: ['Grass', 'Poison']
		},
		{
			id: 4,
			name: 'Charmander',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/4.png',
			types: ['Fire']
		},
		{
			id: 7,
			name: 'Squirtle',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/7.png',
			types: ['Water']
		},
		{
			id: 25,
			name: 'Pikachu',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/25.png',
			types: ['Electric']
		},
		{
			id: 94,
			name: 'Gengar',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/94.png',
			types: ['Ghost', 'Poison']
		},
		{
			id: 150,
			name: 'Mewtwo',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/150.png',
			types: ['Psychic']
		},
		{
			id: 6,
			name: 'Charizard',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/6.png',
			types: ['Fire', 'Flying']
		},
		{
			id: 149,
			name: 'Dragonite',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/149.png',
			types: ['Dragon', 'Flying']
		},
		{
			id: 384,
			name: 'Rayquaza',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/384.png',
			types: ['Dragon', 'Flying']
		}
	];
</script>

<div class="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
	{#each pokemonList as pokemon}
		<div
			class="overflow-hidden rounded-lg bg-white shadow-md transition-shadow duration-300 hover:shadow-lg"
		>
			<div class="relative p-4">
				<div class="absolute right-2 top-2 text-xs font-semibold text-gray-500">
					#{pokemon.id.toString().padStart(3, '0')}
				</div>
				<div class="flex justify-center pb-2 pt-4">
					<img
						src={pokemon.image}
						alt={pokemon.name}
						class="h-32 w-32 transform object-contain transition-transform duration-300 hover:scale-110"
					/>
				</div>
				<h3 class="mb-2 text-center text-lg font-semibold text-gray-800">
					{pokemon.name}
				</h3>
				<div class="flex justify-center space-x-2">
					{#each pokemon.types as type}
						<span
							class={`rounded-full px-2 py-1 text-xs font-medium text-white ${typeColors[type]}`}
						>
							{type}
						</span>
					{/each}
				</div>
			</div>
			<div class="border-t border-gray-100 p-4">
				<a
					href={`/pokemon/${pokemon.id}`}
					class="block rounded-lg bg-red-500 py-2 text-center text-white transition-colors duration-300 hover:bg-red-600"
				>
					View Details
				</a>
			</div>
		</div>
	{/each}
</div>
