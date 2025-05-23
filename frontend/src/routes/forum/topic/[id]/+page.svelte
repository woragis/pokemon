<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { MessageSquare, Heart, Eye } from 'lucide-svelte';
	import type { ForumTopicResponse, ForumComment } from '$lib/types/forum';
	import {
		fetchForumTopicById,
		fetchForumCommentsByTopicId,
		likeForumTopicById,
		createForumComment
	} from '$lib/api/forum';

	let topic: ForumTopicResponse | null = null;
	let comments: ForumComment[] = [];
	let loading = true;
	let error: string | null = null;
	let liked = false;

	let newComment = '';
	let postingComment = false;

	// Pagination state
	let limit = 10;
	let offset = 0;
	let hasMore = true;
	let loadingMore = false;

	$: topicId = $page.params.id;

	onMount(async () => {
		try {
			topic = await fetchForumTopicById({ id: topicId });
			if (topic) {
				await loadComments();
				await fetchForumTopicById({ id: topicId });
				topic.views += 1;
			}
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
		} finally {
			loading = false;
		}
	});

	async function loadComments() {
		if (!hasMore || loadingMore) return;

		loadingMore = true;
		try {
			const res = await fetchForumCommentsByTopicId({ id: topicId, limit, offset });
			comments.push(...res.comments);
			offset += res.comments.length;
			hasMore = res.comments.length === limit;
		} catch (err) {
			console.error('Failed to load comments:', err);
			hasMore = false;
		} finally {
			loadingMore = false;
		}
	}

	async function toggleLike() {
		if (!topic) return;
		try {
			await likeForumTopicById({ id: topicId });
			topic.likes += liked ? 1 : -1;
		} catch (err) {
			console.error('Failed to like topic:', err);
		}
	}

	async function submitComment() {
		if (!newComment.trim()) return;
		postingComment = true;
		try {
			if (topic) {
				const comment = await createForumComment({ id: topicId, content: newComment });
				comments = [...comments, comment];
				topic.replies += 1;
				newComment = '';
			}
		} catch (err) {
			alert('Failed to post comment');
		} finally {
			postingComment = false;
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

		<div class="mb-6 flex space-x-6 text-sm text-gray-600">
			<div class="flex items-center">
				<MessageSquare class="mr-1" size={16} />
				{topic.replies} replies
			</div>
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

		<!-- New Comment Form -->
		<div class="mb-4">
			<textarea
				class="w-full rounded border p-2"
				placeholder="Write a comment..."
				bind:value={newComment}
				rows={3}
			></textarea>
			<button
				class="mt-2 rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700 disabled:opacity-50"
				disabled={postingComment}
				on:click={submitComment}
			>
				{postingComment ? 'Posting...' : 'Post Comment'}
			</button>
		</div>

		<!-- Comments List -->
		<div class="space-y-4">
			{#each comments as comment}
				<div class="border-t pt-4">
					<div class="mb-2 flex items-center text-sm text-gray-600">
						<img src={comment.user.avatar} alt="" class="mr-2 h-6 w-6 rounded-full" />
						<span>{comment.user.username} • {comment.created_at}</span>
					</div>
					<p class="text-gray-800">{comment.content}</p>
				</div>
			{/each}

			{#if hasMore}
				<button
					class="mt-4 w-full rounded bg-gray-100 py-2 text-sm hover:bg-gray-200 disabled:opacity-50"
					disabled={loadingMore}
					on:click={loadComments}
				>
					{loadingMore ? 'Loading more...' : 'Load More Comments'}
				</button>
			{/if}
		</div>
	</div>
{/if}
