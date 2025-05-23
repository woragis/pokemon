<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { putForumTopic } from '$lib/api/forum';

	let title = '';
	let content = '';
	let categoryId = '';
	let categories: { id: string; name: string }[] = [];
	let loading = true;
	let topicId = '';

	$: topicId = $page.params.id;

	onMount(async () => {
		try {
			const resCategories = await fetch('/api/forum/categories');
			categories = await resCategories.json();

			const resTopic = await fetch(`/api/forum/topics/${topicId}`);
			if (!resTopic.ok) throw new Error('Failed to load topic');

			const topic = await resTopic.json();
			title = topic.title;
			content = topic.content;
			categoryId = topic.categoryId;
		} catch (err) {
			alert('Failed to load topic or categories');
		} finally {
			loading = false;
		}
	});

	async function submit() {
		const res = await putForumTopic({
			id: topicId,
			topic: {
				title,
				content,
				category_id: categoryId
			}
		});

		if (res.ok) {
			goto(`/forum/topic/${topicId}`);
		} else {
			alert('Error updating topic');
		}
	}
</script>

{#if loading}
	<p class="p-6 text-gray-500">Loading topic...</p>
{:else}
	<div class="mx-auto max-w-xl p-6">
		<h2 class="mb-4 text-xl font-bold">Edit Topic</h2>

		<form on:submit|preventDefault={submit} class="space-y-4">
			<div>
				<label for="title" class="mb-1 block text-sm font-medium">Title</label>
				<input
					id="title"
					type="text"
					bind:value={title}
					class="w-full rounded border px-3 py-2"
					required
				/>
			</div>

			<div>
				<label for="category" class="mb-1 block text-sm font-medium">Category</label>
				<select
					id="category"
					bind:value={categoryId}
					class="w-full rounded border px-3 py-2"
					required
				>
					<option disabled value="">Select category</option>
					{#each categories as cat}
						<option value={cat.id}>{cat.name}</option>
					{/each}
				</select>
			</div>

			<div>
				<label for="content" class="mb-1 block text-sm font-medium">Content</label>
				<textarea
					id="content"
					bind:value={content}
					rows="6"
					class="w-full rounded border px-3 py-2"
					required
				></textarea>
			</div>

			<button type="submit" class="rounded bg-blue-600 px-4 py-2 text-white">Update Topic</button>
		</form>
	</div>
{/if}
