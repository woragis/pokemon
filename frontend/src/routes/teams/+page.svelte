<script>
	import TeamBuilder from '$lib/components/teams/TeamBuilder.svelte';
	import { Shield, AlertTriangle } from 'lucide-svelte';
</script>

<div class="bg-gradient-to-r from-blue-600 to-blue-700 px-4 py-16">
	<div class="container mx-auto">
		<h1 class="mb-4 text-center text-3xl font-bold text-white md:text-4xl">Team Builder</h1>
		<p class="mx-auto mb-6 max-w-2xl text-center text-blue-100">
			Create the perfect Pokémon team. Analyze type advantages, build balanced teams, and prepare
			for battles.
		</p>
	</div>
</div>

<div class="container mx-auto px-4 py-12">
	<div class="grid grid-cols-1 gap-8 lg:grid-cols-4">
		<div class="lg:col-span-3">
			<TeamBuilder />
		</div>

		<div class="lg:col-span-1">
			<div class="mb-6 rounded-lg bg-white p-6 shadow-md">
				<h3 class="mb-4 text-lg font-semibold">Team Analysis</h3>

				<div class="mb-4">
					<div class="mb-2 flex items-center justify-between">
						<span class="text-sm font-medium">Offensive Coverage</span>
						<span class="text-sm font-medium text-green-600">Strong</span>
					</div>
					<div class="h-2 w-full rounded-full bg-gray-200">
						<div class="h-2 rounded-full bg-green-500" style="width: 75%"></div>
					</div>
				</div>

				<div class="mb-4">
					<div class="mb-2 flex items-center justify-between">
						<span class="text-sm font-medium">Defensive Coverage</span>
						<span class="text-sm font-medium text-yellow-600">Average</span>
					</div>
					<div class="h-2 w-full rounded-full bg-gray-200">
						<div class="h-2 rounded-full bg-yellow-500" style="width: 50%"></div>
					</div>
				</div>

				<div class="mb-4">
					<div class="mb-2 flex items-center justify-between">
						<span class="text-sm font-medium">Type Diversity</span>
						<span class="text-sm font-medium text-green-600">Excellent</span>
					</div>
					<div class="h-2 w-full rounded-full bg-gray-200">
						<div class="h-2 rounded-full bg-green-500" style="width: 90%"></div>
					</div>
				</div>

				<div class="mt-6">
					<h4 class="mb-2 flex items-center text-sm font-semibold">
						<Shield size="16" class="mr-1 text-blue-600" />
						Weaknesses
					</h4>
					<div class="text-sm text-gray-700">
						<p class="mb-1">Your team has weaknesses against:</p>
						<ul class="list-disc pl-5 text-sm">
							<li>Ground-type attacks</li>
							<li>Ice-type attacks</li>
						</ul>
					</div>
				</div>

				<div class="mt-4">
					<h4 class="mb-2 flex items-center text-sm font-semibold">
						<AlertTriangle size="16" class="mr-1 text-yellow-600" />
						Suggestions
					</h4>
					<div class="text-sm text-gray-700">
						<p>Consider adding a Pokémon with resistance to Ground and Ice attacks.</p>
					</div>
				</div>
			</div>

			<div class="rounded-lg bg-white p-6 shadow-md">
				<h3 class="mb-4 text-lg font-semibold">Saved Teams</h3>
				<div class="space-y-3">
					<div class="rounded-md bg-gray-50 p-3">
						<h4 class="font-medium">Elite Four Challenge</h4>
						<p class="mt-1 text-xs text-gray-500">Last updated: 2 days ago</p>
					</div>
					<div class="rounded-md bg-gray-50 p-3">
						<h4 class="font-medium">Water Type Team</h4>
						<p class="mt-1 text-xs text-gray-500">Last updated: 1 week ago</p>
					</div>
					<div class="rounded-md bg-gray-50 p-3">
						<h4 class="font-medium">Championship Team</h4>
						<p class="mt-1 text-xs text-gray-500">Last updated: 2 weeks ago</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
