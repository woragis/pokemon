<script>
	import FeatureCard from './FeatureCard.svelte';
	import { BookOpen, Users, Newspaper, MessageSquare } from 'lucide-svelte';

	const features = [
		{
			title: 'Complete Pokédex',
			description:
				'Track your Pokémon collection with our comprehensive Pokédex featuring all generations and regional variants.',
			icon: BookOpen,
			color: 'bg-gradient-to-br from-red-500 to-red-600',
			link: '/pokedex'
		},
		{
			title: 'Team Builder',
			description:
				'Create the perfect team with our advanced team builder. Analyze strengths, weaknesses, and compatible Pokémon.',
			icon: Users,
			color: 'bg-gradient-to-br from-blue-500 to-blue-600',
			link: '/teams'
		},
		{
			title: 'Latest News',
			description:
				'Stay updated with the latest Pokémon news, game updates, events, and community highlights.',
			icon: Newspaper,
			color: 'bg-gradient-to-br from-yellow-500 to-yellow-600',
			link: '/news'
		},
		{
			title: 'Community Forum',
			description:
				'Connect with trainers worldwide, share strategies, ask questions, and join discussions in our active community.',
			icon: MessageSquare,
			color: 'bg-gradient-to-br from-green-500 to-green-600',
			link: '/forum'
		}
	];
</script>

<section class="feature-section">
	<div class="feature-container">
		<div class="feature-header">
			<h2 class="feature-title">Feature-Packed Pokémon Hub</h2>
			<p class="feature-subtitle">
				Everything you need to enhance your Pokémon journey in one place.
			</p>
		</div>

		<div class="feature-grid">
			{#each features as feature, index}
				<FeatureCard
					title={feature.title}
					description={feature.description}
					color={feature.color}
					link={feature.link}
					icon={feature.icon}
				/>
			{/each}
		</div>
	</div>
</section>

<style>
	.feature-section {
		background-color: white;
		padding: 80px 0;
	}

	.feature-container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 0 1rem;
	}

	.feature-header {
		margin-bottom: 64px;
		text-align: center;
	}

	.feature-title {
		font-size: 2rem;
		font-weight: bold;
		color: #1f2937; /* gray-900 */
		margin-bottom: 1rem;
	}

	@media (min-width: 768px) {
		.feature-title {
			font-size: 2.5rem;
		}
	}

	.feature-subtitle {
		max-width: 640px;
		margin: 0 auto;
		font-size: 1.25rem;
		color: #4b5563; /* gray-600 */
	}

	.feature-grid {
		display: grid;
		grid-template-columns: 1fr;
		gap: 2rem;
	}

	@media (min-width: 768px) {
		.feature-grid {
			grid-template-columns: repeat(2, 1fr);
		}
	}

	@media (min-width: 1024px) {
		.feature-grid {
			grid-template-columns: repeat(4, 1fr);
		}
	}
</style>
