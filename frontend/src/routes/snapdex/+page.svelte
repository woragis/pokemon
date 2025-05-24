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
	<h1 class="title">Snapdex Feed</h1>

	{#if error}
		<p class="error">{error}</p>
	{:else if snaps.length === 0}
		<p class="empty">No snaps yet. Be the first!</p>
	{:else}
		<ul class="feed-grid">
			{#each snaps as snap}
				<li>
					<FeedPicture {snap} />
				</li>
			{/each}
		</ul>
	{/if}

	<div class="feed-links">
		<a href="/snapdex/create" class="feed-link">New</a>
		<a href="/snapdex/me" class="feed-link">Me</a>
	</div>
</section>

<style>
	.snapdex-feed {
		padding: 2rem 1rem;
		max-width: 1200px;
		margin: 0 auto;
	}

	.snapdex-feed .title {
		font-size: 2rem;
		font-weight: bold;
		text-align: center;
		color: #1e40af; /* Blue-800 */
		margin-bottom: 1.5rem;
	}

	.snapdex-feed .error {
		color: #dc2626; /* Red-600 */
		text-align: center;
	}

	.snapdex-feed .empty {
		text-align: center;
		color: #64748b; /* Gray-500 */
	}

	.feed-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
		gap: 1rem;
	}

	.feed-picture {
		display: block;
		border-radius: 0.5rem;
		overflow: hidden;
		background-color: #ffffff;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
		transition: transform 0.2s ease;
		text-decoration: none;
		color: #111827; /* Gray-900 */
	}

	.feed-picture:hover {
		transform: scale(1.02);
	}

	.feed-picture img {
		width: 100%;
		height: 200px;
		object-fit: cover;
		display: block;
	}

	.feed-picture p {
		padding: 0.5rem 0.75rem;
		font-size: 0.95rem;
		color: #1e293b; /* Gray-800 */
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	.feed-links {
		margin-top: 2rem;
		display: flex;
		justify-content: center;
		gap: 1rem;
	}

	.feed-link {
		background-color: #2563eb; /* Blue-600 */
		color: white;
		padding: 0.5rem 1.25rem;
		border-radius: 0.5rem;
		text-decoration: none;
		font-weight: 600;
		transition: background-color 0.2s ease;
	}

	.feed-link:hover {
		background-color: #1d4ed8; /* Blue-700 */
	}
</style>
