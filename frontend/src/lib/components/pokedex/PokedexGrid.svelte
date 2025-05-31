<script lang="ts">
	import { fetchAllPokemons } from '$lib/api/pokedex';
	import { pagination, pokemons } from '$lib/store/pokemons';
	import { onMount } from 'svelte';
	import PokedexPokemon from './PokedexPokemon.svelte';

	let loading = false;

	$: {
		const { offset, limit } = $pagination;
		loading = true;
		fetchAllPokemons({ offset, limit }).then(() => {
			loading = false;
		});
	}

	onMount(() => {
		fetchAllPokemons($pagination);
	});

	export function getPokemonIdFromUrl(url: string): number {
		const match = url.match(/\/pokemon\/(\d+)\//);
		if (!match) throw new Error('Invalid Pokémon URL');
		return parseInt(match[1], 10);
	}
</script>

<div class="pokemon-grid">
	{#if loading}
		<h1>Loading</h1>
	{:else}
		{#each $pokemons as pokemon}
			<PokedexPokemon {pokemon} />
		{/each}
	{/if}
</div>

<style>
	.pokemon-grid {
		display: grid;
		grid-template-columns: 1fr;
		gap: 1.5rem; /* Tailwind's gap-6 (6 * 0.25rem) = 1.5rem */
	}

	@media (min-width: 640px) {
		/* sm: ≥ 640px */
		.pokemon-grid {
			grid-template-columns: repeat(2, 1fr);
		}
	}

	@media (min-width: 768px) {
		/* md: ≥ 768px */
		.pokemon-grid {
			grid-template-columns: repeat(3, 1fr);
		}
	}

	@media (min-width: 1024px) {
		/* lg: ≥ 1024px */
		.pokemon-grid {
			grid-template-columns: repeat(4, 1fr);
		}
	}
</style>
