<script lang="ts">
	interface NewsItem {
		id: number;
		title: string;
		summary: string;
		date: string;
		image: string;
		category: string;
	}

	const newsItems: NewsItem[] = [
		{
			id: 1,
			title: 'New Pokémon Game Announced for 2025',
			summary:
				'Game Freak reveals new open-world Pokémon adventure coming to the latest Nintendo console.',
			date: 'June 15, 2025',
			image:
				'https://images.pexels.com/photos/371924/pexels-photo-371924.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260',
			category: 'Games'
		},
		{
			id: 2,
			title: 'Pokémon World Championships 2025 Dates Announced',
			summary: 'The biggest competitive Pokémon event is coming to Tokyo this August.',
			date: 'May 28, 2025',
			image:
				'https://images.pexels.com/photos/2116719/pexels-photo-2116719.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260',
			category: 'Events'
		},
		{
			id: 3,
			title: 'New Legendary Pokémon Revealed in Upcoming Movie',
			summary: 'The latest Pokémon movie teases a never-before-seen legendary creature.',
			date: 'May 12, 2025',
			image:
				'https://images.pexels.com/photos/3800507/pexels-photo-3800507.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260',
			category: 'Movies'
		},
		{
			id: 4,
			title: 'Major Balance Update Coming to Pokémon Mobile Games',
			summary:
				'Niantic announces significant changes to Pokémon types and abilities in upcoming update.',
			date: 'April 30, 2025',
			image:
				'https://images.pexels.com/photos/5082579/pexels-photo-5082579.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260',
			category: 'Mobile'
		}
	];

	const categoryColors: Record<string, string> = {
		Games: 'bg-blue-500',
		Events: 'bg-green-500',
		Movies: 'bg-red-500',
		Mobile: 'bg-purple-500',
		Anime: 'bg-yellow-500',
		Merchandise: 'bg-pink-500'
	};
</script>

<div class="space-y-8">
	{#each newsItems as item}
		<div
			class="overflow-hidden rounded-lg bg-white shadow-md transition-shadow duration-300 hover:shadow-lg"
		>
			<div class="md:flex">
				<div class="md:w-1/3">
					<img src={item.image} alt={item.title} class="h-48 w-full object-cover md:h-full" />
				</div>

				<div class="p-6 md:w-2/3">
					<div class="mb-3 flex items-center">
						<span
							class={`rounded-full px-2 py-1 text-xs font-medium text-white ${categoryColors[item.category]}`}
						>
							{item.category}
						</span>
						<span class="ml-3 text-sm text-gray-500">{item.date}</span>
					</div>

					<h3 class="mb-3 text-xl font-bold text-gray-900">{item.title}</h3>
					<p class="mb-4 text-gray-600">{item.summary}</p>

					<a
						href={`/news/${item.id}`}
						class="inline-block font-semibold text-red-600 transition-colors hover:text-red-800"
					>
						Read More →
					</a>
				</div>
			</div>
		</div>
	{/each}
</div>
