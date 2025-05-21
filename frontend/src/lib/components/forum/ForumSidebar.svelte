<script>
	import { MessageCircle } from 'lucide-svelte';
</script>

<div class="lg:col-span-1">
	<div class="mb-6 rounded-lg bg-white p-6 shadow-md">
		<h3 class="mb-4 text-lg font-semibold">Categories</h3>
		<ul class="space-y-2">
			{#each [{ name: 'General Discussion', count: 38 }, { name: 'Competitive Battles', count: 24 }, { name: 'Trading', count: 16 }, { name: 'Game Strategies', count: 29 }, { name: 'Fan Creations', count: 12 }] as category}
				<li>
					<a href="#" class="flex items-center justify-between text-gray-700 hover:text-green-600">
						<span>{category.name}</span>
						<span class="rounded-full bg-gray-100 px-2 py-1 text-xs text-gray-600"
							>{category.count}</span
						>
					</a>
				</li>
			{/each}
		</ul>
	</div>

	<div class="rounded-lg bg-white p-6 shadow-md">
		<h3 class="mb-4 flex items-center text-lg font-semibold">
			<MessageCircle size={18} class="mr-2 text-green-600" />
			Active Discussions
		</h3>
		<div class="space-y-4">
			<a href="#" class="group block">
				<h4 class="font-medium transition-colors group-hover:text-green-600">
					Which starter Pokémon is best for speedruns?
				</h4>
				<p class="mt-1 text-sm text-gray-500">32 replies • 2 hours ago</p>
			</a>
			<a href="#" class="group block">
				<h4 class="font-medium transition-colors group-hover:text-green-600">
					Need help with Elite Four team composition
				</h4>
				<p class="mt-1 text-sm text-gray-500">19 replies • 5 hours ago</p>
			</a>
			<a href="#" class="group block">
				<h4 class="font-medium transition-colors group-hover:text-green-600">
					Share your favorite Pokémon designs!
				</h4>
				<p class="mt-1 text-sm text-gray-500">47 replies • 1 day ago</p>
			</a>
		</div>
	</div>
</div>
