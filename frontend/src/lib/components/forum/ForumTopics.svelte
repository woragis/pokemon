<script lang="ts">
	import { onMount } from 'svelte';
	import { MessageSquare, Heart, Eye } from 'lucide-svelte';
	import type { ForumTopicResponse } from '$lib/types/forum';
	import { fetchForumTopics } from '$lib/api/forum';

	let forumTopics: ForumTopicResponse[] = [];
	let loading = true;
	let error: string | null = null;

	const categoryColors: Record<string, string> = {
		Competitive: 'bg-red-500',
		General: 'bg-blue-500',
		Strategies: 'bg-green-500',
		Collecting: 'bg-purple-500',
		Trading: 'bg-yellow-500',
		Events: 'bg-pink-500'
	};

	onMount(async () => {
		try {
			forumTopics = await fetchForumTopics();
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			loading = false;
		}
	});
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

	{#if loading}
		<p class="px-6 py-4 text-gray-500">Loading...</p>
	{:else if error}
		<p class="px-6 py-4 text-red-500">Error: {error}</p>
	{:else if forumTopics.length === 0}
		<p class="px-6 py-4 text-gray-500">No topics found.</p>
	{:else}
		<div class="divide-y">
			{#each forumTopics as topic}
				<div class={`transition-colors hover:bg-gray-50 ${topic.pinned ? 'bg-yellow-50' : ''}`}>
					<div class="px-6 py-4">
						<div class="md:flex md:items-center md:justify-between">
							<div class="md:flex-1">
								<div class="mb-2 flex items-center">
									<span
										class={`rounded-full px-2 py-1 text-xs font-medium text-white ${categoryColors[topic.category.color] ?? 'bg-gray-400'}`}
									>
										{topic.category.name}
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
										alt={topic.author.avatar}
										class="mr-2 h-6 w-6 rounded-full"
									/>
									<span class="text-sm text-gray-600">
										<span class="font-medium text-gray-900">{topic.author.username}</span> • {topic.created_at}
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
	{/if}
</div>
