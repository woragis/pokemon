<script>
	import { Search, ChevronRight } from 'lucide-svelte';
</script>

<div class="relative overflow-hidden bg-gradient-to-br from-red-600 to-red-700">
	<!-- Decorative Pokeball pattern in the background -->
	<div class="absolute inset-0 z-0 opacity-10">
		<div class="absolute left-10 top-10 h-40 w-40 rounded-full border-8 border-white"></div>
		<div class="absolute bottom-20 right-20 h-60 w-60 rounded-full border-8 border-white"></div>
		<div class="absolute right-40 top-40 h-20 w-20 rounded-full border-4 border-white"></div>
		<div class="w-30 h-30 absolute bottom-40 left-40 rounded-full border-4 border-white"></div>
	</div>

	<div class="container relative z-10 mx-auto px-4 py-20 md:py-32">
		<div class="mx-auto max-w-3xl text-center">
			<h1 class="mb-6 text-4xl font-bold leading-tight text-white md:text-6xl">
				Your Ultimate Pokémon Companion
			</h1>
			<p class="mb-10 text-xl text-red-100 md:text-2xl">
				Track your Pokédex progress, build perfect teams, and connect with trainers worldwide.
			</p>

			<div class="mb-10 flex flex-col justify-center gap-4 sm:flex-row">
				<a
					href="/pokedex"
					class="flex items-center justify-center gap-2 rounded-full bg-yellow-400 px-8 py-3 font-semibold text-gray-900 transition-colors hover:bg-yellow-500"
				>
					Explore Pokédex <ChevronRight size={18} />
				</a>
				<a
					href="/teams"
					class="flex items-center justify-center gap-2 rounded-full bg-blue-600 px-8 py-3 font-semibold text-white transition-colors hover:bg-blue-700"
				>
					Build a Team <ChevronRight size={18} />
				</a>
			</div>

			<div class="relative mx-auto max-w-md">
				<div class="relative">
					<input
						type="text"
						placeholder="Search for a Pokémon..."
						class="w-full rounded-full border-2 border-white/30 bg-white/20 px-5 py-3 pr-12 text-white placeholder-white/70 backdrop-blur-sm focus:outline-none focus:ring-2 focus:ring-yellow-400"
					/>
					<div class="absolute right-3 top-1/2 -translate-y-1/2 transform">
						<Search class="h-5 w-5 text-white" />
					</div>
				</div>
			</div>
		</div>
	</div>

	<!-- Wave decoration at the bottom -->
	<div class="absolute bottom-0 left-0 right-0">
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 100" class="h-auto w-full">
			<path
				fill="#ffffff"
				fill-opacity="1"
				d="M0,64L48,69.3C96,75,192,85,288,80C384,75,480,53,576,42.7C672,32,768,32,864,42.7C960,53,1056,75,1152,80C1248,85,1344,75,1392,69.3L1440,64L1440,100L1392,100C1344,100,1248,100,1152,100C1056,100,960,100,864,100C768,100,672,100,576,100C480,100,384,100,288,100C192,100,96,100,48,100L0,100Z"
			></path>
		</svg>
	</div>
</div>
