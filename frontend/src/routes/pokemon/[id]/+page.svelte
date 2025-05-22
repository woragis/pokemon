<script lang="ts">
	import { pokemons } from '$lib/store/pokemons';
	import { page } from '$app/stores';
	import { derived } from 'svelte/store';

	const pokemonId = derived(page, ($page) => Number($page.params.id));
	const pokemon = derived([pokemons, pokemonId], ([$pokemons, $id]) =>
		$pokemons.find((p) => p.id === $id)
	);
</script>

{#if $pokemon}
	<div class="container mx-auto px-4 py-10">
		<h1 class="mb-6 text-center text-4xl font-bold capitalize">{$pokemon.name}</h1>

		<div class="flex flex-col gap-10 md:flex-row">
			<!-- Left Column: Image and types -->
			<div class="flex flex-col items-center md:w-1/3">
				<img
					src={$pokemon.sprites.other['official-artwork'].front_default ||
						$pokemon.sprites.front_default}
					alt={$pokemon.name}
					class="mb-4 w-60"
				/>

				<div class="flex flex-wrap gap-2">
					{#each $pokemon.types as t}
						<span class="rounded-full bg-gray-200 px-4 py-1 text-sm capitalize">{t.type.name}</span>
					{/each}
				</div>
			</div>

			<!-- Right Column: Data sections -->
			<div class="md:w-2/3">
				<h2 class="mb-4 text-2xl font-semibold">Pokédex Data</h2>
				<div class="grid grid-cols-2 gap-4 text-sm">
					<p><strong>National №:</strong> {$pokemon.id.toString().padStart(4, '0')}</p>
					<p><strong>Species:</strong> {$pokemon.name} Pokémon</p>
					<p><strong>Height:</strong> {$pokemon.height / 10} m</p>
					<p><strong>Weight:</strong> {$pokemon.weight / 10} kg</p>
					<p><strong>Base Exp:</strong> {$pokemon.base_experience}</p>
					<p><strong>Abilities:</strong></p>
					<ul class="col-span-2 ml-2 list-inside list-disc">
						{#each $pokemon.abilities as ab}
							<li>{ab.ability.name}</li>
						{/each}
					</ul>
				</div>

				<hr class="my-6" />

				<h2 class="mb-4 text-2xl font-semibold">Base Stats</h2>
				<div class="space-y-2">
					{#each $pokemon.stats as stat}
						<div>
							<div class="mb-1 flex justify-between text-sm">
								<span class="capitalize">{stat.stat.name}</span>
								<span>{stat.base_stat}</span>
							</div>
							<div class="h-2 w-full rounded bg-gray-200">
								<div
									class="h-2 rounded bg-red-500"
									style="width: {Math.min(stat.base_stat, 100)}%"
								></div>
							</div>
						</div>
					{/each}
				</div>
			</div>
		</div>
	</div>
{:else}
	<div class="container mx-auto px-4 py-12 text-center text-gray-500">
		<p>Pokémon not found.</p>
	</div>
{/if}
