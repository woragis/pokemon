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
	let element: HTMLElement;
	let visible = false;

	onMount(() => {
		interval = setInterval(() => {
			isAnimating = true;
			setTimeout(() => {
				currentIndex = (currentIndex + 1) % samplePokemon.length;
				isAnimating = false;
			}, 300);
		}, 5000);
		const observer = new IntersectionObserver(
			([entry]) => {
				if (entry.isIntersecting) {
					visible = true;
					observer.disconnect();
				}
			},
			{ threshold: 0.1 }
		);

		if (element) observer.observe(element);
	});

	onDestroy(() => {
		clearInterval(interval);
	});
</script>

<section class="featured-section">
	<div class="featured-container">
		<div class="featured-header">
			<h2 class="featured-title">Featured Pokémon</h2>
			<p class="featured-subtitle">Discover popular Pokémon and learn about their strengths.</p>
		</div>

		<div bind:this={element} class="featured-content" class:animate={visible}>
			<div class={`pokemon-card ${isAnimating ? 'hidden' : 'visible'}`}>
				<div class="pokemon-image-wrapper">
					<img
						src={samplePokemon[currentIndex].image}
						alt={samplePokemon[currentIndex].name}
						class="pokemon-image"
					/>
				</div>

				<div class="pokemon-info">
					<div class="pokemon-id">
						<span>#00{samplePokemon[currentIndex].id}</span>
					</div>

					<h3 class="pokemon-name">{samplePokemon[currentIndex].name}</h3>

					<div class="pokemon-types">
						{#each samplePokemon[currentIndex].types as type}
							<span class={`pokemon-type ${typeColors[type]}`}>{type}</span>
						{/each}
					</div>

					<p class="pokemon-description">
						This {samplePokemon[currentIndex].name} is known for its exceptional abilities and strength
						in battle. Click to learn more about its stats, moves, and evolution chain.
					</p>

					<a href={`/pokemon/${samplePokemon[currentIndex].id}`} class="view-details-button">
						View Details
					</a>
				</div>
			</div>

			<div class="pokemon-pagination">
				{#each samplePokemon as _, index}
					<button
						on:click={() => changePokemon(index)}
						class={`pagination-dot ${index === currentIndex ? 'active' : ''}`}
						aria-label={`View Pokémon ${index + 1}`}
					></button>
				{/each}
			</div>
		</div>
	</div>
</section>

<style>
	.featured-section {
		background: linear-gradient(to bottom right, #2563eb, #1e40af); /* from-blue-600 to-blue-700 */
		padding: 5rem 0;
		overflow: hidden;
	}

	.featured-container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 0 1rem;
	}

	.featured-header {
		text-align: center;
		margin-bottom: 3rem;
	}

	.featured-title {
		font-size: 2rem;
		font-weight: bold;
		color: white;
		margin-bottom: 1rem;
	}

	.featured-subtitle {
		font-size: 1.25rem;
		color: #dbeafe; /* text-blue-100 */
		max-width: 36rem;
		margin: 0 auto;
	}

	.featured-content {
		max-width: 64rem;
		margin: 0 auto;
		transform: translateY(20px);
		opacity: 0;
	}
	@keyframes fadeInUp {
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}
	.featured-content.animate {
		animation: fadeInUp 0.5s ease forwards;
	}

	.pokemon-card {
		display: flex;
		flex-direction: column;
		align-items: center;
		background: white;
		padding: 2rem;
		border-radius: 1rem;
		box-shadow: 0 10px 20px rgba(0, 0, 0, 0.15);
		transition: opacity 0.3s ease;
	}

	.pokemon-card.visible {
		opacity: 1;
	}

	.pokemon-card.hidden {
		opacity: 0;
		pointer-events: none;
	}

	@media (min-width: 768px) {
		.pokemon-card {
			flex-direction: row;
			padding: 3rem;
		}
	}

	.pokemon-image-wrapper {
		margin-bottom: 2rem;
		width: 100%;
		max-width: 16rem;
	}

	.pokemon-image {
		width: 100%;
		transition: transform 0.3s ease;
	}

	.pokemon-image:hover {
		transform: scale(1.05);
	}

	@media (min-width: 768px) {
		.pokemon-image-wrapper {
			margin-right: 3rem;
			margin-bottom: 0;
			width: 50%;
		}
	}

	.pokemon-info {
		text-align: center;
		width: 100%;
	}

	@media (min-width: 768px) {
		.pokemon-info {
			text-align: left;
			width: 50%;
		}
	}

	.pokemon-id {
		margin-bottom: 1rem;
		color: #6b7280; /* gray-500 */
		font-size: 0.875rem;
		font-weight: 500;
	}

	.pokemon-name {
		font-size: 2rem;
		font-weight: bold;
		color: #111827; /* gray-900 */
		margin-bottom: 1rem;
	}

	.pokemon-types {
		display: flex;
		justify-content: center;
		gap: 0.5rem;
		margin-bottom: 1.5rem;
	}

	@media (min-width: 768px) {
		.pokemon-types {
			justify-content: flex-start;
		}
	}

	.pokemon-type {
		padding: 0.25rem 0.75rem;
		border-radius: 9999px;
		font-size: 0.875rem;
		font-weight: 500;
		color: white;
	}

	.pokemon-description {
		color: #4b5563; /* gray-600 */
		margin-bottom: 1.5rem;
	}

	.view-details-button {
		display: inline-block;
		background-color: #2563eb;
		color: white;
		font-weight: 600;
		padding: 0.75rem 1.5rem;
		border-radius: 0.5rem;
		text-decoration: none;
		transition: background-color 0.3s ease;
	}

	.view-details-button:hover {
		background-color: #1e40af;
	}

	.pokemon-pagination {
		display: flex;
		justify-content: center;
		gap: 0.5rem;
		margin-top: 2rem;
	}

	.pagination-dot {
		width: 0.75rem;
		height: 0.75rem;
		border-radius: 9999px;
		background-color: rgba(255, 255, 255, 0.4);
		border: none;
		cursor: pointer;
		transition: background-color 0.3s ease;
	}

	.pagination-dot.active {
		background-color: white;
	}
</style>
