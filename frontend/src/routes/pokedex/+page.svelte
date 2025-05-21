<script>
	import PokedexGrid from '$lib/components/pokemon/PokedexGrid.svelte';
	import { Search } from 'lucide-svelte';

	let pokemonList = [
		{ id: 1, name: 'Bulbasaur' },
		{ id: 4, name: 'Charmander' },
		{ id: 7, name: 'Squirtle' }
	];

	// You can add reactive variables and handlers here later for search input, filtering, pagination, etc.
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

	<PokedexGrid />

	<div class="mt-12 flex justify-center">
		<nav class="flex items-center">
			<button
				class="flex h-10 w-10 items-center justify-center rounded-l-lg border border-gray-300 text-gray-600 hover:bg-gray-100 disabled:cursor-not-allowed disabled:opacity-50"
				disabled
			>
				Prev
			</button>
			<button
				class="flex h-10 w-10 items-center justify-center border-b border-t border-gray-300 bg-red-600 font-medium text-white"
			>
				1
			</button>
			<button
				class="flex h-10 w-10 items-center justify-center border-b border-t border-gray-300 text-gray-600 hover:bg-gray-100"
			>
				2
			</button>
			<button
				class="flex h-10 w-10 items-center justify-center border-b border-t border-gray-300 text-gray-600 hover:bg-gray-100"
			>
				3
			</button>
			<button
				class="flex h-10 w-10 items-center justify-center rounded-r-lg border border-gray-300 text-gray-600 hover:bg-gray-100"
			>
				Next
			</button>
		</nav>
	</div>
</div>
