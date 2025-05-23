<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { MessageSquare, Heart, Eye } from 'lucide-svelte';
	import type { ForumTopicResponse } from '$lib/types/forum';
	import { fetchForumTopicById } from '$lib/api/forum';

	let topic: ForumTopicResponse | null = null;
	let loading = true;
	let error: string | null = null;

	$: topicId = $page.params.id;

	onMount(async () => {
		try {
			topic = await fetchForumTopicById({ id: topicId });
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			loading = false;
		}
	});
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
			<!-- Placeholder for actual content -->
			<p>{topic.content}</p>
		</div>

		<div class="flex space-x-6 text-sm text-gray-600">
			<div class="flex items-center">
				<MessageSquare class="mr-1" size={16} />
				{topic.replies} replies
			</div>
			<div class="flex items-center"><Heart class="mr-1" size={16} /> {topic.likes} likes</div>
			<div class="flex items-center"><Eye class="mr-1" size={16} /> {topic.views} views</div>
		</div>
	</div>
{/if}
