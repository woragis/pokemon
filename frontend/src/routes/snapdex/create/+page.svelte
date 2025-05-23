<script lang="ts">
	import { createSnap, type Snap } from '$lib/api/snapdex';
	import { goto } from '$app/navigation';

	let media_url = '';
	let caption = '';
	let tags = '';
	let error: string | null = null;

	async function handleSubmit() {
		error = null;
		try {
			const newSnap = await createSnap({
				media_url,
				caption,
				tags: tags.split(',').map((tag) => tag.trim())
			});
			goto(`/snapdex/${newSnap.id}`);
		} catch {
			error = 'Failed to post snap.';
		}
	}
</script>

<h1>Create a Snap</h1>

<form on:submit|preventDefault={handleSubmit} class="space-y-4">
	{#if error}
		<p class="text-red-600">{error}</p>
	{/if}
	<input bind:value={media_url} placeholder="Media URL" class="input" required />
	<input bind:value={caption} placeholder="Caption" class="input" required />
	<input bind:value={tags} placeholder="Tags (comma-separated)" class="input" />

	<button type="submit" class="btn btn-primary">Post Snap</button>
</form>
