<script lang="ts">
	import { deleteForumCategory, fetchForumCategories, postForumCategory } from '$lib/api/forum';
	import type { ForumCategory } from '$lib/types/forum';
	import { onMount } from 'svelte';

	let name = '';
	let color = '';
	let description = '';
	let categories: ForumCategory[] = [];

	onMount(async () => {
		await loadCategories();
	});

	async function loadCategories() {
		categories = await fetchForumCategories();
	}

	async function createCategory() {
		const res = await postForumCategory({ category: { name, color, description } });
		if (res.ok) {
			name = '';
			color: '';
			description = '';
			await loadCategories();
		}
	}

	async function deleteCategory(id: string) {
		await deleteForumCategory({ id });
		await loadCategories();
	}
</script>

<div class="mx-auto max-w-2xl p-6">
	<h2 class="mb-4 text-xl font-bold">Manage Categories</h2>

	<form on:submit|preventDefault={createCategory} class="mb-6 space-y-4">
		<input
			type="text"
			placeholder="Category name"
			bind:value={name}
			class="w-full rounded border px-3 py-2"
			required
		/>
		<input
			type="text"
			placeholder="Category description"
			bind:value={description}
			class="w-full rounded border px-3 py-2"
			required
		/>
		<input
			type="text"
			placeholder="Category color (use tailwind colors ex.: bg-red-500)"
			bind:value={color}
			class="w-full rounded border px-3 py-2"
			required
		/>
		<button class="rounded bg-green-600 px-4 py-2 text-white">Create Category</button>
	</form>

	<ul class="space-y-2">
		{#each categories as cat}
			<li class="flex items-center justify-between rounded border px-4 py-2">
				<div class="flex flex-col">
					<div class="flex">
						<span class="font-semibold">{cat.name}</span>
						<span class={`ml-2 rounded px-2 py-1 text-xs text-white ${cat.color}`}>{cat.color}</span
						>
					</div>
					<span class="ml-2 rounded px-2 py-1 text-xl">{cat.description}</span>
				</div>
				<button on:click={() => deleteCategory(cat.id)} class="text-red-500">Delete</button>
			</li>
		{/each}
	</ul>
</div>
