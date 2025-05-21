<script lang="ts">
	import { fetchShoutById } from '$lib/api/shouts';
	import type { Shout } from '$lib/types/shouts';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	let shout: Shout;
	$: id = $page.params.id;

	onMount(async () => {
		shout = await fetchShoutById({ id });
	});
</script>

{#if shout}
	<div class="mx-auto max-w-2xl rounded border p-4 shadow">
		<h1 class="mb-2 text-xl font-semibold">{shout.content}</h1>
		<p class="text-sm text-gray-500">
			By {shout.author} — {new Date(shout.created_at).toLocaleString()}
		</p>
		<div class="mt-4">
			<a href="/shouts" class="text-blue-600">← Back</a>
		</div>
	</div>
{:else}
	<p class="mt-8 text-center">Loading...</p>
{/if}
