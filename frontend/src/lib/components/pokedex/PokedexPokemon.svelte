<script lang="ts">
	import { typeColors } from '$lib/store/pokemons';
	import type { PokemonData } from '$lib/types/pokeapi';
	import { onMount } from 'svelte';

	const imgPrefix =
		'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/';

	export function capitalize(str: string): string {
		return str.charAt(0).toUpperCase() + str.slice(1);
	}
	export let pokemon: PokemonData;

	let element: HTMLElement;
	let visible = false;

	onMount(() => {
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
</script>

<div bind:this={element} class="pokemon-card" class:animate={visible}>
	<div class="card-content">
		<div class="card-id">#{pokemon.id}</div>
		<div class="pokemon-image-wrapper">
			<img
				src={`${imgPrefix}/${pokemon.id}.png`}
				alt={pokemon.name}
				class="pokemon-image"
				loading="lazy"
			/>
		</div>
		<h3 class="pokemon-name">{capitalize(pokemon.name)}</h3>
		<div class="pokemon-types">
			{#each pokemon.types as type}
				<span class="pokemon-type {typeColors[type.type.name]}">
					{capitalize(type.type.name)}
				</span>
			{/each}
		</div>
	</div>
	<div class="card-footer">
		<a href={`/pokemon/${pokemon.id}`} class="details-button"> View Details </a>
	</div>
</div>

<style>
	@keyframes fadeInUp {
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.pokemon-card {
		background-color: #ffffff;
		border-radius: 0.75rem;
		overflow: hidden;
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
		transition: box-shadow 0.3s ease;
		display: flex;
		flex-direction: column;
		opacity: 0;
		transform: translateY(-20px);
	}
	.pokemon-card.animate {
		animation: fadeInUp 0.5s ease forwards;
	}

	.pokemon-card:hover {
		box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
	}

	.card-content {
		position: relative;
		padding: 1rem;
		flex: 1;
	}

	.card-id {
		position: absolute;
		top: 0.5rem;
		right: 0.5rem;
		font-size: 0.75rem;
		font-weight: 600;
		color: #6b7280; /* gray-500 */
	}

	.pokemon-image-wrapper {
		display: flex;
		justify-content: center;
		padding: 1rem 0 0.5rem;
	}

	.pokemon-image {
		height: 8rem;
		width: 8rem;
		object-fit: contain;
		transition: transform 0.3s ease;
	}

	.pokemon-image:hover {
		transform: scale(1.1);
	}

	.pokemon-name {
		text-align: center;
		font-size: 1.125rem; /* ~text-lg */
		font-weight: 600;
		color: #1f2937; /* gray-800 */
		margin-bottom: 0.5rem;
	}

	.pokemon-types {
		display: flex;
		justify-content: center;
		gap: 0.5rem;
		flex-wrap: wrap;
	}

	.pokemon-type {
		border-radius: 9999px;
		padding: 0.25rem 0.5rem;
		font-size: 0.75rem;
		font-weight: 500;
		color: #fff;
	}

	.card-footer {
		border-top: 1px solid #f3f4f6; /* gray-100 */
		padding: 1rem;
	}

	.details-button {
		display: block;
		text-align: center;
		background-color: #ef4444; /* red-500 */
		color: #ffffff;
		padding: 0.5rem;
		border-radius: 0.5rem;
		text-decoration: none;
		font-weight: 500;
		transition: background-color 0.3s ease;
	}

	.details-button:hover {
		background-color: #dc2626; /* red-600 */
	}
</style>
