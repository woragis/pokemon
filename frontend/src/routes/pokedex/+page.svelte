<script lang="ts">
	import { fetchAllPokemons } from '$lib/api/pokedex';
	import type { Pagination } from '$lib/api/pokedex/get';
	import PokedexGrid from '$lib/components/pokemon/PokedexGrid.svelte';
	import { pokemons, total } from '$lib/store/pokemons';
	import { Search } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	const pagination = writable<Pagination>({ offset: 0, limit: 20 });

	function nextPage() {
		pagination.update((p) => {
			const next = p.offset + p.limit;
			return { ...p, offset: next };
		});
		loadPage();
	}

	function prevPage() {
		pagination.update((p) => {
			const prev = Math.max(p.offset - p.limit, 0);
			return { ...p, offset: prev };
		});
		loadPage();
	}

	const loadPage = async () => {
		const $pagination = get(pagination);
		await fetchAllPokemons($pagination);
	};

	import { get } from 'svelte/store';

	onMount(() => {
		loadPage();
	});
</script>

<div class="bg-gradient-to-r from-red-600 to-red-700 px-4 py-16">
	<div class="container mx-auto">
		<h1 class="mb-4 text-center text-3xl font-bold text-white md:text-4xl">Pokédex</h1>
		<p class="mx-auto mb-8 max-w-2xl text-center text-red-100">
			Explore the complete Pokémon database. Track your collection, learn about different species,
			and discover their abilities.
		</p>

		<div class="relative mx-auto max-w-xl">
			<div class="relative">
				<input
					type="text"
					placeholder="Search Pokémon by name or number..."
					class="w-full rounded-full border-2 border-white/30 bg-white/20 px-5 py-3 pr-12 text-white placeholder-white/70 backdrop-blur-sm focus:outline-none focus:ring-2 focus:ring-yellow-400"
				/>
				<div class="absolute right-3 top-1/2 -translate-y-1/2 transform">
					<Search class="h-5 w-5 text-white" />
				</div>
			</div>
		</div>
	</div>
</div>

<div class="container mx-auto px-4 py-12">
	<div class="mb-8 rounded-lg bg-gray-100 p-4">
		<div class="flex flex-wrap gap-2">
			<button class="rounded-full bg-red-600 px-4 py-2 text-sm text-white">All Types</button>
			<button class="rounded-full border border-gray-300 bg-white px-4 py-2 text-sm text-gray-700"
				>Fire</button
			>
			<button class="rounded-full border border-gray-300 bg-white px-4 py-2 text-sm text-gray-700"
				>Water</button
			>
			<button class="rounded-full border border-gray-300 bg-white px-4 py-2 text-sm text-gray-700"
				>Grass</button
			>
			<button class="rounded-full border border-gray-300 bg-white px-4 py-2 text-sm text-gray-700"
				>Electric</button
			>
			<button class="rounded-full border border-gray-300 bg-white px-4 py-2 text-sm text-gray-700"
				>Psychic</button
			>
			<button class="rounded-full border border-gray-300 bg-white px-4 py-2 text-sm text-gray-700"
				>Dragon</button
			>
		</div>
	</div>

	<PokedexGrid pokemons={$pokemons} />

	<div class="mt-12 flex justify-center">
		<nav class="flex items-center gap-2">
			<button
				on:click={prevPage}
				class="flex items-center gap-1 rounded-full bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm ring-1 ring-gray-300 transition hover:bg-red-50 disabled:cursor-not-allowed disabled:opacity-50"
				disabled={$pagination.offset === 0}
			>
				<span>←</span>
				<span>Previous</span>
			</button>

			<span class="rounded-full bg-red-600 px-4 py-2 text-sm font-bold text-white shadow">
				Page {$pagination.offset / $pagination.limit + 1}
			</span>

			<button
				on:click={nextPage}
				class="flex items-center gap-1 rounded-full bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm ring-1 ring-gray-300 transition hover:bg-red-50 disabled:cursor-not-allowed disabled:opacity-50"
				disabled={$pagination.offset + $pagination.limit >= $total}
			>
				<span>Next</span>
				<span>→</span>
			</button>
		</nav>
	</div>
</div>
