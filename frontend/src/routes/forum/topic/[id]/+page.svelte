<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { MessageSquare, Heart, Eye } from 'lucide-svelte';
	import type { ForumTopicResponse } from '$lib/types/forum';
	import { fetchForumTopicById, likeForumTopicById, viewForumTopicById } from '$lib/api/forum';

	let topic: ForumTopicResponse | null = null;
	let loading = true;
	let error: string | null = null;
	let liked = false;

	$: topicId = $page.params.id;

	onMount(async () => {
		try {
			topic = await fetchForumTopicById({ id: topicId });

			if (topic) {
				await viewForumTopicById({ id: topicId });
				topic.views += 1;
			}
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			loading = false;
		}
	});

	async function toggleLike() {
		if (!topic) return;
		try {
			const change = await likeForumTopicById({ id: topicId });
			liked = change.liked;
			topic.likes += liked ? 1 : -1;
		} catch (err) {
			console.error('Failed to like topic:', err);
		}
	}
</script>

{#if loading}
	<p class="p-6 text-gray-500">Loading topic...</p>
{:else if error}
	<p class="p-6 text-red-500">Error: {error}</p>
{:else if topic}
	<div class="mx-auto max-w-3xl p-6">
		<h1 class="mb-2 text-2xl font-bold">{topic.title}</h1>

		<div class="mb-4 flex items-center text-sm text-gray-500">
			<img src={topic.author.avatar} alt={''} class="mr-2 h-6 w-6 rounded-full" />
			<span>{topic.author.username} • {topic.created_at} • {topic.category.name}</span>
		</div>

		<div class="prose mb-6">
			<p>{topic.content}</p>
		</div>

		<div class="flex space-x-6 text-sm text-gray-600">
			<div class="flex items-center">
				<MessageSquare class="mr-1" size={16} />
				{topic.replies} replies
			</div>

			<!-- Like Button -->
			<button class="flex items-center hover:text-red-600" on:click={toggleLike}>
				<Heart class="mr-1" size={16} fill={liked ? 'currentColor' : 'none'} />
				{topic.likes}
				{liked ? 'liked' : 'likes'}
			</button>

			<div class="flex items-center">
				<Eye class="mr-1" size={16} />
				{topic.views} views
			</div>
		</div>
	</div>
{/if}
