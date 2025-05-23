<script lang="ts">
	import { getMySnaps, type Snap } from '$lib/api/snapdex';
	import { onMount } from 'svelte';

	let snaps: Snap[] = [];
	let error: string | null = null;

	onMount(async () => {
		try {
			snaps = await getMySnaps();
		} catch {
			error = 'Could not load your snaps';
		}
	});
</script>

<h1>My Snaps</h1>

{#if error}
	<p class="text-red-600">{error}</p>
{:else if snaps.length === 0}
	<p>You haven’t posted anything yet.</p>
{:else}
	<ul class="grid grid-cols-2 gap-4 md:grid-cols-3">
		{#each snaps as snap}
			<li>
				<a href={`/snapdex/${snap.id}`}>
					<img src={snap.media_url} alt="Snap" class="h-40 w-full rounded object-cover" />
					<p>{snap.caption}</p>
				</a>
			</li>
		{/each}
	</ul>
{/if}
