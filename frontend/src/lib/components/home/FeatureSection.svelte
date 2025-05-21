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

<section class="bg-white py-20">
	<div class="container mx-auto px-4">
		<div class="mb-16 text-center">
			<h2 class="mb-4 text-3xl font-bold text-gray-900 md:text-4xl">Feature-Packed Pokémon Hub</h2>
			<p class="mx-auto max-w-2xl text-xl text-gray-600">
				Everything you need to enhance your Pokémon journey in one place.
			</p>
		</div>

		<div class="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-4">
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
