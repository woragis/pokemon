<script lang="ts">
	import { goto } from '$app/navigation';
	import { fetchForumCategories, postForumTopic } from '$lib/api/forum';
	import type { ForumCategory } from '$lib/types/forum';
	import { onMount } from 'svelte';

	let title = '';
	let content = '';
	let categoryId = '';
	let categories: ForumCategory[] = [];

	onMount(async () => {
		categories = await fetchForumCategories();
	});

	async function submit() {
		const res = await postForumTopic({ topic: { title, content, category_id: categoryId } });

		if (res.ok) {
			const topic = await res.json();
			goto(`/forum/topic/${topic.id}`);
		} else {
			alert('Error creating topic');
		}
	}
</script>

<div class="mx-auto max-w-xl p-6">
	<h2 class="mb-4 text-xl font-bold">Create a New Topic</h2>

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

		<button type="submit" class="rounded bg-blue-600 px-4 py-2 text-white">Post Topic</button>
	</form>
</div>
