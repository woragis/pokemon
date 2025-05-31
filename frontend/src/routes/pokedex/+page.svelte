<script lang="ts">
	import PokedexNav from '$lib/components/pokedex/PokedexNav.svelte';
	import PokedexGrid from '$lib/components/pokedex/PokedexGrid.svelte';
	import { Search } from 'lucide-svelte';
</script>

<div class="hero-section">
	<div class="container">
		<h1 class="hero-title">Pokédex</h1>
		<p class="hero-subtitle">
			Explore the complete Pokémon database. Track your collection, learn about different species,
			and discover their abilities.
		</p>

		<div class="search-wrapper">
			<div class="search-box">
				<input type="text" placeholder="Search Pokémon by name or number..." class="search-input" />
				<div class="search-icon">
					<Search class="search-icon-inner" />
				</div>
			</div>
		</div>
	</div>
</div>

<div class="main-content">
	<div class="filter-bar">
		<div class="filter-buttons">
			<button class="filter-button active">All Types</button>
			<button class="filter-button">Fire</button>
			<button class="filter-button">Water</button>
			<button class="filter-button">Grass</button>
			<button class="filter-button">Electric</button>
			<button class="filter-button">Psychic</button>
			<button class="filter-button">Dragon</button>
		</div>
	</div>

	<PokedexGrid />
	<PokedexNav />
</div>

<style>
	/* HERO SECTION */

	.hero-section {
		background: linear-gradient(to right, #dc2626, #b91c1c); /* from-red-600 to-red-700 */
		padding: 4rem 1rem;
		color: white;
	}

	.container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 0 1rem;
	}

	.hero-title {
		text-align: center;
		font-size: 2rem;
		font-weight: bold;
		margin-bottom: 1rem;
	}

	@media (min-width: 768px) {
		.hero-title {
			font-size: 2.5rem;
		}
	}

	.hero-subtitle {
		text-align: center;
		color: #fecaca; /* red-100 */
		max-width: 40rem;
		margin: 0 auto 2rem;
		font-size: 1rem;
	}

	/* SEARCH */

	.search-wrapper {
		max-width: 36rem;
		margin: 0 auto;
		position: relative;
	}

	.search-box {
		position: relative;
	}

	.search-input {
		width: 100%;
		padding: 0.75rem 3rem 0.75rem 1.25rem;
		border-radius: 9999px;
		border: 2px solid rgba(255, 255, 255, 0.3);
		background-color: rgba(255, 255, 255, 0.2);
		backdrop-filter: blur(6px);
		color: white;
		font-size: 1rem;
		placeholder-color: rgba(255, 255, 255, 0.7);
	}

	.search-input::placeholder {
		color: rgba(255, 255, 255, 0.7);
	}

	.search-input:focus {
		outline: none;
		box-shadow: 0 0 0 3px #facc15; /* ring-yellow-400 */
	}

	.search-icon {
		position: absolute;
		top: 50%;
		right: 0.75rem;
		transform: translateY(-50%);
		pointer-events: none;
	}

	.search-icon-inner {
		width: 1.25rem;
		height: 1.25rem;
		color: white;
	}

	/* MAIN CONTENT */

	.main-content {
		max-width: 1200px;
		margin: 0 auto;
		padding: 3rem 1rem;
	}

	.filter-bar {
		background-color: #f3f4f6; /* gray-100 */
		border-radius: 0.5rem;
		padding: 1rem;
		margin-bottom: 2rem;
	}

	.filter-buttons {
		display: flex;
		flex-wrap: wrap;
		gap: 0.5rem;
	}

	.filter-button {
		border: 1px solid #d1d5db; /* gray-300 */
		background-color: white;
		color: #374151; /* gray-700 */
		font-size: 0.875rem;
		padding: 0.5rem 1rem;
		border-radius: 9999px;
		cursor: pointer;
		transition: background-color 0.3s ease;
	}

	.filter-button:hover {
		background-color: #f3f4f6;
	}

	.filter-button.active {
		background-color: #dc2626; /* red-600 */
		color: white;
		border: none;
	}
</style>
