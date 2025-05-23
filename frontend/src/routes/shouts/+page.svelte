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

<div class="mx-auto max-w-2xl px-4 pt-6">
	<!-- Header -->
	<div class="mb-6 border-b pb-4">
		<h1 class="text-3xl font-semibold text-gray-900">Shouts</h1>
	</div>

	<!-- New Shout Button -->
	<div class="mb-6">
		<a
			href="/shouts/new"
			class="inline-flex items-center rounded-full bg-blue-500 px-6 py-2 font-medium text-white transition hover:bg-blue-600"
		>
			➕ New Shout
		</a>
	</div>

	<!-- Shout Feed -->
	{#if shouts.length === 0}
		<p class="text-center text-gray-500">No shouts yet.</p>
	{:else}
		<ul class="space-y-6">
			{#each shouts as shout}
				<li
					class="rounded-xl border border-gray-200 bg-white p-5 shadow-sm transition hover:shadow-md"
				>
					<!-- Content -->
					<p class="whitespace-pre-line text-base text-gray-800">{shout.content}</p>

					<!-- Meta Info -->
					<p class="mt-3 text-sm text-gray-500">
						<span class="font-semibold text-gray-700">{shout.user.name}</span>
						<span class="mx-1">·</span>
						{new Date(shout.created_at).toLocaleString()}
					</p>

					<!-- Action Buttons -->
					<div class="mt-4 flex gap-4 text-sm text-gray-600">
						<a href={`/shouts/${shout.id}`} class="text-blue-600 hover:underline">View</a>
						<a href={`/shouts/${shout.id}/edit`} class="text-yellow-600 hover:underline">Edit</a>
						<button on:click={() => handleDelete(shout.id)} class="text-red-600 hover:underline">
							Delete
						</button>
					</div>
				</li>
			{/each}
		</ul>
	{/if}
</div>
