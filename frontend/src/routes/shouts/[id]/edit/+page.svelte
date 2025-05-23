<script lang="ts">
	import { fetchShoutById, putShout } from '$lib/api/shouts';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import type { Shout } from '$lib/types/shouts';

	let shout: Shout;
	let content = '';
	let author = '';
	$: id = $page.params.id;

	onMount(async () => {
		shout = await fetchShoutById({ id });
		content = shout.content;
		author = shout.author;
	});

	const submit = async () => {
		await putShout({ id, shout: { content } });
		goto(`/shouts/${id}`);
	};
</script>

<div class="mx-auto max-w-md p-4">
	<h1 class="mb-4 text-2xl font-bold">Edit Shout</h1>
	<form on:submit|preventDefault={submit} class="space-y-4">
		<textarea bind:value={content} class="w-full rounded border p-2" rows="4" required></textarea>
		<button type="submit" class="rounded bg-yellow-500 px-4 py-2 text-white">Update</button>
	</form>
</div>
