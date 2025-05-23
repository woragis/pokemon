<script lang="ts">
	import { page } from '$app/stores';
	import { getSnap, type Snap } from '$lib/api/snapdex';
	import { onMount } from 'svelte';

	let snap: Snap | null = null;
	let error: string | null = null;

	$: id = $page.params.id;

	onMount(async () => {
		try {
			snap = await getSnap(id);
		} catch {
			error = 'Failed to load snap';
		}
	});
</script>

{#if error}
	<p class="text-red-600">{error}</p>
{:else if !snap}
	<p>Loading snap...</p>
{:else}
	<div class="space-y-4">
		<img src={snap.media_url} alt="Snap" class="w-full rounded-md" />
		<h2 class="text-xl font-semibold">{snap.caption}</h2>
		<p class="text-gray-500">Tags: {snap.tags.join(', ')}</p>
		<p class="text-sm text-gray-400">Posted on {new Date(snap.created_at).toLocaleDateString()}</p>
	</div>
{/if}
