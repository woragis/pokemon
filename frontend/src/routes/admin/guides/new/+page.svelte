<script>
	import { createGuide } from '@/api/guides.js';

	let title = '';
	let category = '';
	let content = '';

	function submit() {
		createGuide({ title, category, content });
		window.location.href = '/admin/guides';
	}
</script>

<h1 class="text-xl font-bold mb-4">New Guide</h1>

<form class="space-y-4" on:submit|preventDefault={submit}>
	<input class="input" bind:value={title} placeholder="Title" />
	<input class="input" bind:value={category} placeholder="Category" />
	<textarea class="input h-40" bind:value={content} placeholder="Content"></textarea>

	<button class="bg-blue-600 text-white px-4 py-2 rounded">Save</button>
</form>
