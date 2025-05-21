<script lang="ts">
	import { Plus, X, Save } from 'lucide-svelte';

	interface Pokemon {
		id: number;
		name: string;
		image: string;
		types: string[];
	}

	let teamName = 'My Team';
	let teamPokemon: Pokemon[] = [];
	let isSelectingPokemon = false;

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

	const availablePokemon: Pokemon[] = [
		{
			id: 25,
			name: 'Pikachu',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/25.png',
			types: ['Electric']
		},
		{
			id: 6,
			name: 'Charizard',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/6.png',
			types: ['Fire', 'Flying']
		},
		{
			id: 9,
			name: 'Blastoise',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/9.png',
			types: ['Water']
		},
		{
			id: 3,
			name: 'Venusaur',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/3.png',
			types: ['Grass', 'Poison']
		},
		{
			id: 149,
			name: 'Dragonite',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/149.png',
			types: ['Dragon', 'Flying']
		},
		{
			id: 94,
			name: 'Gengar',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/94.png',
			types: ['Ghost', 'Poison']
		}
	];

	function addPokemonToTeam(pokemon: Pokemon) {
		if (teamPokemon.length < 6) {
			teamPokemon = [...teamPokemon, pokemon];
			isSelectingPokemon = false;
		}
	}

	function removePokemonFromTeam(index: number) {
		teamPokemon = [...teamPokemon.slice(0, index), ...teamPokemon.slice(index + 1)];
	}

	function saveTeam() {
		alert(`Team "${teamName}" saved with ${teamPokemon.length} Pokémon!`);
	}
</script>

<div class="mx-auto max-w-4xl">
	<div class="mb-8 rounded-lg bg-white p-6 shadow-md">
		<div class="mb-6 flex flex-col justify-between md:flex-row md:items-center">
			<div class="mb-4 md:mb-0">
				<label for="team-name" class="mb-1 block text-sm font-medium text-gray-700">Team Name</label
				>
				<input
					id="team-name"
					bind:value={teamName}
					class="w-full rounded-md border border-gray-300 px-4 py-2 focus:border-blue-500 focus:ring-blue-500 md:w-64"
				/>
			</div>

			<button
				on:click={saveTeam}
				class="flex items-center justify-center rounded-md px-4 py-2 font-medium text-white"
				class:bg-gray-400={teamPokemon.length === 0}
				class:cursor-not-allowed={teamPokemon.length === 0}
				class:bg-green-600={teamPokemon.length > 0}
				class:hover:bg-green-700={teamPokemon.length > 0}
				disabled={teamPokemon.length === 0}
			>
				<Save size={18} class="mr-2" />
				Save Team
			</button>
		</div>

		<div class="grid grid-cols-2 gap-4 md:grid-cols-3 lg:grid-cols-6">
			{#each Array(6) as _, index}
				<div
					class="flex h-40 flex-col items-center justify-center rounded-lg border-2 {teamPokemon[
						index
					]
						? 'border-gray-200'
						: 'border-dashed border-gray-300'}"
				>
					{#if teamPokemon[index]}
						<div class="relative flex h-full w-full flex-col items-center justify-center p-2">
							<button
								on:click={() => removePokemonFromTeam(index)}
								class="absolute right-1 top-1 flex h-6 w-6 items-center justify-center rounded-full bg-red-500 text-white"
								aria-label="Remove Pokemon"
							>
								<X size={14} />
							</button>
							<img
								src={teamPokemon[index].image}
								alt={teamPokemon[index].name}
								class="h-20 w-20 object-contain"
							/>
							<h3 class="mt-1 text-center text-sm font-medium">{teamPokemon[index].name}</h3>
							<div class="mt-1 flex justify-center space-x-1">
								{#each teamPokemon[index].types as type}
									<span
										class={`rounded-full px-2 py-0.5 text-xs font-medium text-white ${typeColors[type]}`}
									>
										{type}
									</span>
								{/each}
							</div>
						</div>
					{:else}
						<button
							on:click={() => (isSelectingPokemon = true)}
							class="flex h-10 w-10 items-center justify-center rounded-full bg-gray-200 hover:bg-gray-300"
							aria-label="Add Pokemon"
						>
							<Plus size={20} class="text-gray-600" />
						</button>
					{/if}
				</div>
			{/each}
		</div>
	</div>

	{#if isSelectingPokemon}
		<div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
			<div class="max-h-[80vh] w-full max-w-2xl overflow-y-auto rounded-lg bg-white p-6">
				<div class="mb-4 flex items-center justify-between">
					<h3 class="text-xl font-semibold">Select a Pokémon</h3>
					<button
						on:click={() => (isSelectingPokemon = false)}
						class="rounded-full p-1 hover:bg-gray-200"
					>
						<X size={20} />
					</button>
				</div>

				<div class="grid grid-cols-2 gap-4 sm:grid-cols-3">
					{#each availablePokemon as pokemon}
						<div
							on:click={() => addPokemonToTeam(pokemon)}
							class="flex cursor-pointer flex-col items-center rounded-lg border p-4 hover:bg-gray-50"
						>
							<img src={pokemon.image} alt={pokemon.name} class="h-24 w-24 object-contain" />
							<h4 class="mt-2 font-medium">{pokemon.name}</h4>
							<div class="mt-1 flex justify-center space-x-1">
								{#each pokemon.types as type}
									<span
										class={`rounded-full px-2 py-0.5 text-xs font-medium text-white ${typeColors[type]}`}
									>
										{type}
									</span>
								{/each}
							</div>
						</div>
					{/each}
				</div>
			</div>
		</div>
	{/if}
</div>
