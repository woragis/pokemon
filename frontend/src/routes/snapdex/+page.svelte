<script lang="ts">
	import { onMount } from 'svelte';
	import type { Snap } from '$lib/api/snapdex';
	import { getAllSnaps } from '$lib/api/snapdex';

	let snaps: Snap[] = [];
	let error: string | null = null;

	onMount(async () => {
		try {
			snaps = await getAllSnaps();
		} catch (e) {
			error = 'Failed to load snaps.';
		}
	});
</script>

<h1>Snapdex Feed</h1>

{#if error}
	<p class="text-red-600">{error}</p>
{:else if snaps.length === 0}
	<p>No snaps yet. Be the first!</p>
{:else}
	<ul class="grid grid-cols-2 gap-4 md:grid-cols-3">
		{#each snaps as snap}
			<li>
				<a href={`/snapdex/${snap.id}`}>
					<img src={snap.media_url} alt="Snap image" class="h-48 w-full rounded object-cover" />
					<p>{snap.caption}</p>
				</a>
			</li>
		{/each}
	</ul>
{/if}
<a href="/snapdex/create">New</a>
<a href="/snapdex/me">Me</a>
