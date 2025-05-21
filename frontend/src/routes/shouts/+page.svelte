<script lang="ts">
	import { onMount } from 'svelte';
	import { fetchShouts, deleteShout } from '$lib/api/shouts';
	import type { Shout } from '$lib/types/shouts';
	import { goto } from '$app/navigation';

	let shouts: Shout[] = [];

	onMount(async () => {
		shouts = await fetchShouts();
	});

	const handleDelete = async (id: string) => {
		await deleteShout({ id });
		shouts = shouts.filter((s) => s.id !== id);
	};
</script>

<div class="mx-auto max-w-3xl p-4">
	<h1 class="mb-4 text-2xl font-bold">Shouts</h1>
	<a href="/shouts/new" class="mb-4 inline-block rounded bg-blue-500 px-4 py-2 text-white"
		>New Shout</a
	>

	{#if shouts.length === 0}
		<p>No shouts yet.</p>
	{:else}
		<ul class="space-y-4">
			{#each shouts as shout}
				<li class="rounded border p-4 shadow-sm">
					<p class="text-lg">{shout.content}</p>
					<p class="mt-1 text-sm text-gray-500">
						By {shout.author} — {new Date(shout.created_at).toLocaleString()}
					</p>
					<div class="mt-2 flex gap-2">
						<a href={`/shouts/${shout.id}`} class="text-blue-600">View</a>
						<a href={`/shouts/${shout.id}/edit`} class="text-yellow-600">Edit</a>
						<button on:click={() => handleDelete(shout.id)} class="text-red-600">Delete</button>
					</div>
				</li>
			{/each}
		</ul>
	{/if}
</div>
