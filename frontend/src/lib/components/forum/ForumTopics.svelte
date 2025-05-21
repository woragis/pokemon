<script lang="ts">
	import { MessageSquare, Heart, Eye } from 'lucide-svelte';

	type ForumTopic = {
		id: number;
		title: string;
		author: string;
		authorAvatar: string;
		date: string;
		replies: number;
		likes: number;
		views: number;
		category: string;
		pinned: boolean;
	};

	const forumTopics: ForumTopic[] = [
		{
			id: 1,
			title: 'Best Fire-type Pokémon for competitive battles?',
			author: 'AshKetchum',
			authorAvatar: 'https://i.pravatar.cc/150?img=1',
			date: '2 hours ago',
			replies: 24,
			likes: 12,
			views: 156,
			category: 'Competitive',
			pinned: true
		},
		{
			id: 2,
			title: "Let's share our favorite Pokémon memories from the games",
			author: 'PokéFan22',
			authorAvatar: 'https://i.pravatar.cc/150?img=2',
			date: '1 day ago',
			replies: 47,
			likes: 35,
			views: 302,
			category: 'General',
			pinned: false
		},
		{
			id: 3,
			title: 'Evolution strategies for Eevee - which one is best?',
			author: 'EeveeLover',
			authorAvatar: 'https://i.pravatar.cc/150?img=3',
			date: '3 days ago',
			replies: 31,
			likes: 18,
			views: 215,
			category: 'Strategies',
			pinned: false
		},
		{
			id: 4,
			title: 'Legendary hunt! Post your rarest catches here',
			author: 'MasterTrainer',
			authorAvatar: 'https://i.pravatar.cc/150?img=4',
			date: '5 days ago',
			replies: 59,
			likes: 42,
			views: 487,
			category: 'Collecting',
			pinned: false
		}
	];

	const categoryColors: Record<string, string> = {
		Competitive: 'bg-red-500',
		General: 'bg-blue-500',
		Strategies: 'bg-green-500',
		Collecting: 'bg-purple-500',
		Trading: 'bg-yellow-500',
		Events: 'bg-pink-500'
	};
</script>

<div class="overflow-hidden rounded-lg bg-white shadow-md">
	<div class="border-b">
		<div class="flex items-center justify-between bg-gray-50 px-6 py-3">
			<div class="flex-grow">
				<h3 class="font-semibold text-gray-900">Topics</h3>
			</div>
			<div class="hidden items-center space-x-8 text-sm text-gray-500 md:flex">
				<span class="w-16 text-center">Replies</span>
				<span class="w-16 text-center">Likes</span>
				<span class="w-16 text-center">Views</span>
			</div>
		</div>
	</div>

	<div class="divide-y">
		{#each forumTopics as topic}
			<div class={`transition-colors hover:bg-gray-50 ${topic.pinned ? 'bg-yellow-50' : ''}`}>
				<div class="px-6 py-4">
					<div class="md:flex md:items-center md:justify-between">
						<div class="md:flex-1">
							<div class="mb-2 flex items-center">
								<span
									class={`rounded-full px-2 py-1 text-xs font-medium text-white ${categoryColors[topic.category]}`}
								>
									{topic.category}
								</span>
								{#if topic.pinned}
									<span
										class="ml-2 rounded-full bg-yellow-100 px-2 py-1 text-xs font-medium text-yellow-800"
									>
										Pinned
									</span>
								{/if}
							</div>

							<a
								href={`/forum/topic/${topic.id}`}
								class="text-lg font-semibold text-gray-900 transition-colors hover:text-red-600"
							>
								{topic.title}
							</a>

							<div class="mt-2 flex items-center">
								<img
									src={topic.authorAvatar}
									alt={topic.author}
									class="mr-2 h-6 w-6 rounded-full"
								/>
								<span class="text-sm text-gray-600">
									<span class="font-medium text-gray-900">{topic.author}</span> • {topic.date}
								</span>
							</div>
						</div>

						<div class="mt-4 flex space-x-4 text-sm text-gray-500 md:mt-0 md:space-x-8">
							<div class="flex items-center md:block">
								<MessageSquare size={16} class="mb-1 mr-1 md:mx-auto md:mr-0" />
								<span class="md:block md:text-center">{topic.replies}</span>
							</div>

							<div class="flex items-center md:block">
								<Heart size={16} class="mb-1 mr-1 md:mx-auto md:mr-0" />
								<span class="md:block md:text-center">{topic.likes}</span>
							</div>

							<div class="flex items-center md:block">
								<Eye size={16} class="mb-1 mr-1 md:mx-auto md:mr-0" />
								<span class="md:block md:text-center">{topic.views}</span>
							</div>
						</div>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>
