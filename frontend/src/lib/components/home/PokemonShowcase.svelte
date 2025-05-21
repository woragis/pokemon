<script lang="ts">
	import { onMount, onDestroy } from 'svelte';

	interface Pokemon {
		id: number;
		name: string;
		image: string;
		types: string[];
	}

	const samplePokemon: Pokemon[] = [
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
			id: 150,
			name: 'Mewtwo',
			image:
				'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/150.png',
			types: ['Psychic']
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

	let currentIndex = 0;
	let isAnimating = false;

	let interval: number | any;

	function changePokemon(index: number) {
		isAnimating = true;
		setTimeout(() => {
			currentIndex = index;
			isAnimating = false;
		}, 300);
	}

	onMount(() => {
		interval = setInterval(() => {
			isAnimating = true;
			setTimeout(() => {
				currentIndex = (currentIndex + 1) % samplePokemon.length;
				isAnimating = false;
			}, 300);
		}, 5000);
	});

	onDestroy(() => {
		clearInterval(interval);
	});
</script>

<section class="overflow-hidden bg-gradient-to-br from-blue-600 to-blue-700 py-20">
	<div class="container mx-auto px-4">
		<div class="mb-12 text-center">
			<h2 class="mb-4 text-3xl font-bold text-white md:text-4xl">Featured Pokémon</h2>
			<p class="mx-auto max-w-2xl text-xl text-blue-100">
				Discover popular Pokémon and learn about their strengths.
			</p>
		</div>

		<div class="mx-auto max-w-4xl">
			<div
				class={`flex flex-col items-center rounded-xl bg-white p-8 shadow-xl transition-opacity duration-300 md:flex-row md:p-12 ${
					isAnimating ? 'opacity-0' : 'opacity-100'
				}`}
			>
				<div class="mb-8 md:mb-0 md:mr-12 md:w-1/2">
					<img
						src={samplePokemon[currentIndex].image}
						alt={samplePokemon[currentIndex].name}
						class="mx-auto w-full max-w-xs transform transition-transform duration-300 hover:scale-105 md:max-w-full"
					/>
				</div>

				<div class="text-center md:w-1/2 md:text-left">
					<div class="mb-4 flex items-center justify-center md:justify-start">
						<span class="text-sm font-medium text-gray-500"
							>#00{samplePokemon[currentIndex].id}</span
						>
					</div>

					<h3 class="mb-4 text-3xl font-bold text-gray-900">{samplePokemon[currentIndex].name}</h3>

					<div class="mb-6 flex items-center justify-center space-x-2 md:justify-start">
						{#each samplePokemon[currentIndex].types as type}
							<span
								class={`rounded-full px-3 py-1 text-sm font-medium text-white ${typeColors[type]}`}
								>{type}</span
							>
						{/each}
					</div>

					<p class="mb-6 text-gray-600">
						This {samplePokemon[currentIndex].name} is known for its exceptional abilities and strength
						in battle. Click to learn more about its stats, moves, and evolution chain.
					</p>

					<a
						href={`/pokemon/${samplePokemon[currentIndex].id}`}
						class="inline-block rounded-lg bg-blue-600 px-6 py-3 font-semibold text-white transition-colors hover:bg-blue-700"
						>View Details</a
					>
				</div>
			</div>

			<div class="mt-8 flex justify-center space-x-2">
				{#each samplePokemon as _, index}
					<button
						on:click={() => changePokemon(index)}
						class={`h-3 w-3 rounded-full transition-colors ${
							index === currentIndex ? 'bg-white' : 'bg-white/40'
						}`}
						aria-label={`View Pokémon ${index + 1}`}
					>
						<!-- aria-label={`View Pokémon ${index + 1}`} -->
					</button>
				{/each}
			</div>
		</div>
	</div>
</section>
