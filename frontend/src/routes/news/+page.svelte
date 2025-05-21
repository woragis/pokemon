<script>
	import NewsList from '$lib/components/news/NewsList.svelte';
</script>

<div class="bg-gradient-to-r from-yellow-500 to-yellow-600 px-4 py-16">
	<div class="container mx-auto">
		<h1 class="mb-4 text-center text-3xl font-bold text-white md:text-4xl">Latest Pokémon News</h1>
		<p class="mx-auto max-w-2xl text-center text-yellow-100">
			Stay updated with the latest Pokémon news, game releases, events, and community announcements.
		</p>
	</div>
</div>

<div class="container mx-auto px-4 py-12">
	<div class="grid grid-cols-1 gap-8 lg:grid-cols-4">
		<div class="lg:col-span-3">
			<NewsList />
		</div>

		<div class="lg:col-span-1">
			<div class="mb-6 rounded-lg bg-white p-6 shadow-md">
				<h3 class="mb-4 text-lg font-semibold">Categories</h3>
				<ul class="space-y-2">
					{#each [{ name: 'Games', count: 12 }, { name: 'Events', count: 8 }, { name: 'Movies', count: 5 }, { name: 'Mobile', count: 9 }, { name: 'Merchandise', count: 7 }] as category}
						<li>
							<a
								href="#"
								class="flex items-center justify-between text-gray-700 hover:text-red-600"
							>
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
				<h3 class="mb-4 text-lg font-semibold">Popular Posts</h3>
				<div class="space-y-4">
					{#each [{ title: 'New Generation Announced at Pokémon Direct', time: '2 weeks ago', views: '1.2k' }, { title: 'Pokémon GO Community Day: Everything You Need to Know', time: '1 month ago', views: '856' }, { title: 'Top 10 Rarest Pokémon Cards Sold in 2025', time: '3 months ago', views: '2.4k' }] as post}
						<a href="#" class="group block">
							<h4 class="font-medium transition-colors group-hover:text-red-600">{post.title}</h4>
							<p class="mt-1 text-sm text-gray-500">{post.time} • {post.views} views</p>
						</a>
					{/each}
				</div>
			</div>
		</div>
	</div>
</div>
