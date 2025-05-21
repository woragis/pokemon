<script lang="ts">
	import { postShout } from '$lib/api/shouts';
	import { goto } from '$app/navigation';

	let content = '';
	let author = '';

	const submit = async () => {
		await postShout({ shout: { content, author } });
		goto('/shouts');
	};
</script>

<div class="mx-auto max-w-md p-4">
	<h1 class="mb-4 text-2xl font-bold">New Shout</h1>
	<form on:submit|preventDefault={submit} class="space-y-4">
		<input bind:value={author} placeholder="Author" class="w-full rounded border p-2" required />
		<textarea
			bind:value={content}
			placeholder="What's happening?"
			class="w-full rounded border p-2"
			rows="4"
			required
		></textarea>
		<button type="submit" class="rounded bg-blue-500 px-4 py-2 text-white">Shout</button>
	</form>
</div>
