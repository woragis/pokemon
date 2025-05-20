<script>
	import { getGuides } from '$lib/api/guides.js';
	let guides = getGuides();
</script>

<h1 class="text-2xl font-bold mb-4">Game Guides</h1>

<ul class="space-y-3">
	{#each guides as g}
		<li class="border p-4 rounded">
			<a href={`/guides/${g.id}`} class="font-semibold text-blue-600">{g.title}</a>
			<p class="text-sm text-gray-500">{g.category}</p>
		</li>
	{/each}
</ul>
