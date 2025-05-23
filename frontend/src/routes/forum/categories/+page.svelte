<script lang="ts">
	import {
		deleteForumCategory,
		fetchForumCategories,
		postForumCategory,
		putForumCategory
	} from '$lib/api/forum';
	import type { ForumCategory } from '$lib/types/forum';
	import { onMount } from 'svelte';

	let name = '';
	let color = '';
	let description = '';
	let categories: ForumCategory[] = [];

	let editingId: string | null = null;

	onMount(async () => {
		await loadCategories();
	});

	async function loadCategories() {
		categories = await fetchForumCategories();
	}

	async function createOrUpdateCategory() {
		if (editingId) {
			await putForumCategory({
				id: editingId,
				category: {
					name,
					color,
					description
				}
			});
		} else {
			await postForumCategory({
				category: {
					name,
					color,
					description
				}
			});
		}

		name = '';
		description = '';
		color = '';
		editingId = null;
		await loadCategories();
	}

	function startEditing(cat: ForumCategory) {
		editingId = cat.id;
		name = cat.name;
		color = cat.color;
		description = cat.description;
	}

	async function deleteCategory(id: string) {
		await deleteForumCategory({ id });
		await loadCategories();
	}
</script>

<div class="mx-auto max-w-2xl p-6">
	<h2 class="mb-4 text-xl font-bold">{editingId ? 'Edit' : 'Create'} Category</h2>

	<form on:submit|preventDefault={createOrUpdateCategory} class="mb-6 space-y-4">
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
			placeholder="Category color (Tailwind color (ex.: bg-red-500))"
			bind:value={color}
			class="w-full rounded border px-3 py-2"
			required
		/>
		<div class="flex gap-2">
			<button class="rounded bg-blue-600 px-4 py-2 text-white"
				>{editingId ? 'Update' : 'Create'}</button
			>
			{#if editingId}
				<button
					type="button"
					on:click={() => {
						name = '';
						color = '';
						editingId = null;
					}}
					class="rounded bg-gray-400 px-4 py-2 text-white"
				>
					Cancel
				</button>
			{/if}
		</div>
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
				<div class="flex gap-2">
					<button on:click={() => startEditing(cat)} class="text-blue-500">Edit</button>
					<button on:click={() => deleteCategory(cat.id)} class="text-red-500">Delete</button>
				</div>
			</li>
		{/each}
	</ul>
</div>
