<script lang="ts">
	import { onMount } from 'svelte';

	type Stat = {
		value: string;
		label: string;
	};

	const stats: Stat[] = [
		{ value: '890+', label: 'Pokémon Species' },
		{ value: '18', label: 'Pokémon Types' },
		{ value: '8', label: 'Generations' },
		{ value: '50K+', label: 'Active Trainers' }
	];

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

<section class="stats-section">
	<div class="stats-container">
		<div bind:this={element} class="stats-grid" class:animate={visible}>
			{#each stats as stat}
				<div class="stat-item">
					<div class="stat-value">{stat.value}</div>
					<div class="stat-label">{stat.label}</div>
				</div>
			{/each}
		</div>
	</div>
</section>

<style>
	@keyframes fadeInUp {
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.stats-section {
		background-color: #f9fafb; /* Tailwind's gray-50 */
		padding-top: 3.5rem; /* 14 * 0.25rem = 3.5rem */
		padding-bottom: 3.5rem;
	}

	.stats-container {
		max-width: 1200px;
		margin-left: auto;
		margin-right: auto;
		padding-left: 1rem; /* 4 * 0.25rem = 1rem */
		padding-right: 1rem;
	}

	.stats-grid {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		gap: 2rem;
		opacity: 0;
		transform: translateY(20px);
	}

	.stats-grid.animate {
		animation: fadeInUp 0.5s ease forwards;
	}

	@media (min-width: 768px) {
		.stats-grid {
			grid-template-columns: repeat(4, 1fr);
		}
	}

	.stat-item {
		text-align: center;
	}

	.stat-value {
		margin-bottom: 0.5rem;
		font-size: 2.25rem; /* Tailwind's text-4xl */
		font-weight: 700;
		color: #dc2626; /* Tailwind's red-600 */
	}

	@media (min-width: 768px) {
		.stat-value {
			font-size: 3rem; /* Tailwind's text-5xl */
		}
	}

	.stat-label {
		color: #4b5563; /* Tailwind's gray-600 */
	}
</style>
