<script lang="ts">
	import type { PokemonData } from '$lib/types/pokeapi';

	export let pokemons: PokemonData[];

	export function getPokemonIdFromUrl(url: string): number {
		const match = url.match(/\/pokemon\/(\d+)\//);
		if (!match) throw new Error('Invalid Pokémon URL');
		return parseInt(match[1], 10);
	}

	const typeColors: Record<string, string> = {
		normal: 'bg-gray-300',
		fire: 'bg-red-500',
		water: 'bg-blue-500',
		electric: 'bg-yellow-400',
		grass: 'bg-green-500',
		ice: 'bg-blue-200',
		fighting: 'bg-red-700',
		poison: 'bg-purple-500',
		ground: 'bg-yellow-600',
		flying: 'bg-indigo-300',
		psychic: 'bg-pink-500',
		bug: 'bg-green-400',
		rock: 'bg-yellow-700',
		ghost: 'bg-purple-700',
		dragon: 'bg-indigo-600',
		dark: 'bg-gray-700',
		steel: 'bg-gray-400',
		fairy: 'bg-pink-300'
	};
	const imgPrefix =
		'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/';

	export function capitalize(str: string): string {
		return str.charAt(0).toUpperCase() + str.slice(1);
	}
</script>

<div class="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
	{#each pokemons as pokemon}
		<div
			class="overflow-hidden rounded-lg bg-white shadow-md transition-shadow duration-300 hover:shadow-lg"
		>
			<div class="relative p-4">
				<div class="absolute right-2 top-2 text-xs font-semibold text-gray-500">
					#{pokemon.id}
				</div>
				<div class="flex justify-center pb-2 pt-4">
					<img
						src={`${imgPrefix}/${pokemon.id}.png`}
						alt={pokemon.name}
						class="h-32 w-32 transform object-contain transition-transform duration-300 hover:scale-110"
					/>
				</div>
				<h3 class="mb-2 text-center text-lg font-semibold text-gray-800">
					{capitalize(pokemon.name)}
				</h3>
				<div class="flex justify-center space-x-2">
					{#each pokemon.types as type}
						<span
							class={`rounded-full px-2 py-1 text-xs font-medium text-white ${typeColors[type.type.name]}`}
						>
							{capitalize(type.type.name)}
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
