<script>
	import { page } from '$app/stores';
	import { getGuide } from '$lib/api/guides.js';

	const guide = getGuide($page.params.id);
</script>

<h1 class="text-2xl font-bold mb-2">{guide.title}</h1>
<p class="text-gray-500 mb-4">{guide.category}</p>

<div class="prose max-w-none">
	{@html guide.content}
</div>
