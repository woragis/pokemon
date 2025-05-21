<script>
	import { Github, Instagram, Twitter } from 'lucide-svelte';
</script>

<footer class="bg-gray-800 pb-8 pt-12 text-white">
	<div class="container mx-auto px-4">
		<div class="grid grid-cols-1 gap-8 md:grid-cols-4">
			<div>
				<div class="mb-4 flex items-center gap-2 text-xl font-bold text-red-400">
					<div class="flex h-6 w-6 items-center justify-center rounded-full bg-red-500">
						<div class="h-2 w-2 rounded-full bg-white"></div>
					</div>
					<span>PokéHub</span>
				</div>
				<p class="mb-4 text-gray-300">
					Your ultimate Pokémon companion for tracking, team building, and community discussions.
				</p>
				<div class="flex space-x-4">
					<a href="#" class="text-gray-300 transition-colors hover:text-white">
						<Github size="20" />
					</a>
					<a href="#" class="text-gray-300 transition-colors hover:text-white">
						<Twitter size="20" />
					</a>
					<a href="#" class="text-gray-300 transition-colors hover:text-white">
						<Instagram size="20" />
					</a>
				</div>
			</div>

			<div>
				<h3 class="mb-4 text-lg font-semibold text-white">Features</h3>
				<ul class="space-y-2">
					<li>
						<a href="/pokedex" class="text-gray-300 transition-colors hover:text-red-300">Pokédex</a
						>
					</li>
					<li>
						<a href="/teams" class="text-gray-300 transition-colors hover:text-red-300"
							>Team Builder</a
						>
					</li>
					<li>
						<a href="/news" class="text-gray-300 transition-colors hover:text-red-300">News</a>
					</li>
					<li>
						<a href="/forum" class="text-gray-300 transition-colors hover:text-red-300"
							>Community Forum</a
						>
					</li>
				</ul>
			</div>

			<div>
				<h3 class="mb-4 text-lg font-semibold text-white">Resources</h3>
				<ul class="space-y-2">
					<li>
						<a href="#" class="text-gray-300 transition-colors hover:text-red-300"
							>API Documentation</a
						>
					</li>
					<li>
						<a href="#" class="text-gray-300 transition-colors hover:text-red-300">Tutorials</a>
					</li>
					<li><a href="#" class="text-gray-300 transition-colors hover:text-red-300">Guides</a></li>
					<li><a href="#" class="text-gray-300 transition-colors hover:text-red-300">FAQ</a></li>
				</ul>
			</div>

			<div>
				<h3 class="mb-4 text-lg font-semibold text-white">Legal</h3>
				<ul class="space-y-2">
					<li>
						<a href="#" class="text-gray-300 transition-colors hover:text-red-300"
							>Terms of Service</a
						>
					</li>
					<li>
						<a href="#" class="text-gray-300 transition-colors hover:text-red-300">Privacy Policy</a
						>
					</li>
					<li>
						<a href="#" class="text-gray-300 transition-colors hover:text-red-300">Cookie Policy</a>
					</li>
					<li>
						<a href="#" class="text-gray-300 transition-colors hover:text-red-300">Copyright</a>
					</li>
				</ul>
			</div>
		</div>

		<div class="mt-8 border-t border-gray-700 pt-8 text-center text-sm text-gray-400">
			<p>Pokémon and its trademarks are ©1995-2025 Nintendo, Creatures, and GAME FREAK.</p>
			<p>
				This is a fan-made application and is not affiliated with Nintendo or The Pokémon Company.
			</p>
			<p class="mt-2">© 2025 PokéHub. All rights reserved.</p>
		</div>
	</div>
</footer>
