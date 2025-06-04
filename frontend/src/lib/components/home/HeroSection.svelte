<script>
	import { ChevronRight, Search } from 'lucide-svelte';
	import pokeball from '$lib/assets/pokeball.png';
	import { onMount } from 'svelte';

	let visible = false;
	onMount(() => {
		visible = true;
	});
</script>

<div class="hero">
	<!-- Decorative Pokeball pattern in the background -->
	<div class="pokeballs">
		<img class:visible class="ball ball-1" src={pokeball} alt="pokeball" />
		<img class:visible class="ball ball-2" src={pokeball} alt="pokeball" />
		<img class:visible class="ball ball-3" src={pokeball} alt="pokeball" />
		<img class:visible class="ball ball-4" src={pokeball} alt="pokeball" />
	</div>

	<div class="content">
		<div class="text-content">
			<h1 class="title">Your Ultimate Pokémon Companion</h1>
			<p class="subtitle">
				Track your Pokédex progress, build perfect teams, and connect with trainers worldwide.
			</p>

			<div class="button-group">
				<a href="/pokedex" class="button yellow">
					Explore Pokédex <ChevronRight size={18} />
				</a>
				<a href="/teams" class="button blue">
					Build a Team <ChevronRight size={18} />
				</a>
			</div>

			<div class="search-wrapper">
				<div class="search-box">
					<input type="text" placeholder="Search for a Pokémon..." class="search-input" />
					<div class="search-icon">
						<Search class="icon" />
					</div>
				</div>
			</div>
		</div>
	</div>

	<div class="wave">
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 100">
			<path
				fill="#ffffff"
				fill-opacity="1"
				d="M0,64L48,69.3C96,75,192,85,288,80C384,75,480,53,576,42.7C672,32,768,32,864,42.7C960,53,1056,75,1152,80C1248,85,1344,75,1392,69.3L1440,64L1440,100L1392,100C1344,100,1248,100,1152,100C1056,100,960,100,864,100C768,100,672,100,576,100C480,100,384,100,288,100C192,100,96,100,48,100L0,100Z"
			></path>
		</svg>
	</div>
</div>

<style>
	.hero {
		position: relative;
		overflow: hidden;
		background: linear-gradient(to bottom right, #dc2626, #d60009);
	}

	.pokeballs {
		position: absolute;
		inset: 0;
		opacity: 0.1;
		z-index: 0;
	}

	.ball {
		position: absolute;
		border-radius: 50%;
		border: none;
		opacity: 0;
		transition: 1s;
		animation: float 8s ease-in-out infinite;
	}

	.ball.visible {
		opacity: 1;
	}

	.ball-1 {
		left: 2.5rem;
		top: 2.5rem;
		width: 10rem;
		height: 10rem;
		border-width: 8px;
		animation-delay: 0s;
	}

	.ball-2 {
		right: 5rem;
		bottom: 5rem;
		width: 15rem;
		height: 15rem;
		border-width: 8px;
		animation-duration: 10s;
		animation-direction: alternate;
	}

	.ball-3 {
		top: 10rem;
		right: 10rem;
		width: 5rem;
		height: 5rem;
		border-width: 4px;
		animation-duration: 5s;
		animation-direction: alternate-reverse;
	}

	.ball-4 {
		bottom: 10rem;
		left: 10rem;
		width: 7rem;
		height: 7rem;
		border-width: 4px;
		animation-duration: 18s;
		animation-direction: reverse;
	}

	@keyframes float {
		0%,
		100% {
			transform: translateY(0px) rotate(-50deg) translateX(0px);
		}
		50% {
			transform: translateY(-20px) rotate(30deg) translateX(-15px);
		}
	}
	.content {
		position: relative;
		z-index: 10;
		max-width: 1280px;
		margin: 0 auto;
		padding: 5rem 1rem;
	}

	@media (min-width: 768px) {
		.content {
			padding: 8rem 1rem;
		}
	}

	.text-content {
		max-width: 768px;
		margin: 0 auto;
		text-align: center;
	}

	.title {
		margin-bottom: 1.5rem;
		font-size: 2.25rem;
		font-weight: bold;
		line-height: 1.2;
		color: white;
	}

	@media (min-width: 768px) {
		.title {
			font-size: 3.75rem;
		}
	}

	.subtitle {
		margin-bottom: 2.5rem;
		font-size: 1.25rem;
		color: #fecaca;
	}

	@media (min-width: 768px) {
		.subtitle {
			font-size: 1.5rem;
		}
	}

	.button-group {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 1rem;
		margin-bottom: 2.5rem;
	}

	@media (min-width: 640px) {
		.button-group {
			flex-direction: row;
		}
	}

	.button {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.5rem;
		padding: 0.75rem 2rem;
		font-weight: 600;
		border-radius: 9999px;
		text-decoration: none;
		transition: background-color 0.3s ease;
	}

	.yellow {
		background-color: #facc15;
		color: #1f2937;
	}

	.yellow:hover {
		background-color: #fbbf24;
	}

	.blue {
		background-color: #2563eb;
		color: white;
	}

	.blue:hover {
		background-color: #1d4ed8;
	}

	.search-wrapper {
		max-width: 28rem;
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
		color: white;
		backdrop-filter: blur(4px);
		outline: none;
	}

	.search-input::placeholder {
		color: rgba(255, 255, 255, 0.7);
	}

	.search-input:focus {
		border-color: #facc15;
		box-shadow: 0 0 0 2px #facc15;
	}

	.search-icon {
		position: absolute;
		right: 0.75rem;
		top: 50%;
		transform: translateY(-50%);
		color: white;
	}

	.icon {
		width: 1.25rem;
		height: 1.25rem;
	}

	.wave {
		position: absolute;
		left: 0;
		right: 0;
		bottom: 0;
	}
</style>
