<script>
	import { deleteGuide, getGuides } from '$lib/api/guides.js';
	let guides = getGuides();
</script>

<h1 class="text-2xl font-bold mb-4">Admin Guide Manager</h1>
<a href="/admin/guides/new" class="text-blue-500">+ New Guide</a>

<ul class="mt-4 space-y-3">
	{#each guides as g}
		<li class="border rounded p-4">
			<div class="flex justify-between items-center">
				<div>
					<h2 class="font-semibold">{g.title}</h2>
					<p class="text-sm text-gray-500">{g.category}</p>
				</div>
				<div class="flex gap-2">
					<a class="text-blue-600" href={`/admin/guides/${g.id}/edit`}>Edit</a>
					<button class="text-red-500" on:click={() => deleteGuide(g.id)}>Delete</button>
				</div>
			</div>
		</li>
	{/each}
</ul>
