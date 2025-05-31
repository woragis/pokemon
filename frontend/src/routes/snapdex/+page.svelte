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
	import FeedPicture from '$lib/components/snapdex/FeedPicture.svelte';
</script>

<section class="snapdex-feed">
	<h1 class="feed-title">Snapdex</h1>

	{#if error}
		<p class="feed-error">{error}</p>
	{:else if snaps.length === 0}
		<p class="feed-empty">No snaps yet. Be the first!</p>
	{:else}
		<ul class="feed-grid">
			{#each snaps as snap}
				<li class="feed-card">
					<FeedPicture {snap} />
				</li>
			{/each}
		</ul>
	{/if}

	<div class="feed-actions">
		<a href="/snapdex/create" class="feed-btn">New Snap</a>
		<a href="/snapdex/me" class="feed-btn secondary">My Snaps</a>
	</div>
</section>

<style>
	.snapdex-feed {
		padding: 2rem 1rem;
		max-width: 1000px;
		margin: 0 auto;
	}

	.feed-title {
		font-size: 2.25rem;
		font-weight: 700;
		text-align: center;
		color: #111827; /* Gray-900 */
		margin-bottom: 2rem;
		font-family: 'Segoe UI', sans-serif;
	}

	.feed-error {
		color: #dc2626; /* Red-600 */
		text-align: center;
		font-size: 1rem;
	}

	.feed-empty {
		color: #6b7280; /* Gray-500 */
		text-align: center;
		font-size: 1rem;
	}

	.feed-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
		gap: 1rem;
	}

	/* Simulating Instagram-like card container */
	.feed-card {
		border-radius: 0.75rem;
		overflow: hidden;
		background: white;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
		transition: transform 0.2s ease;
		cursor: pointer;
	}

	.feed-card:hover {
		transform: scale(1.01);
	}

	/* Assume FeedPicture handles internal styling (e.g., image, captions) */

	.feed-actions {
		margin-top: 2rem;
		display: flex;
		justify-content: center;
		gap: 1rem;
		flex-wrap: wrap;
	}

	.feed-btn {
		background-color: #000;
		color: white;
		padding: 0.6rem 1.5rem;
		border-radius: 0.5rem;
		text-decoration: none;
		font-weight: 600;
		font-size: 0.95rem;
		transition:
			background-color 0.2s ease,
			transform 0.1s ease;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
	}

	.feed-btn:hover {
		background-color: #111;
		transform: translateY(-1px);
	}

	.feed-btn.secondary {
		background-color: #f3f4f6; /* Gray-100 */
		color: #111827;
	}

	.feed-btn.secondary:hover {
		background-color: #e5e7eb; /* Gray-200 */
	}
</style>
