<script lang="ts">
	import ForumTopics from '$lib/components/forum/ForumTopics.svelte';
	import ForumSidebar from '$lib/components/forum/ForumSidebar.svelte';
	import ForumControls from '$lib/components/forum/ForumControls.svelte';
	import { PlusCircle, Settings } from 'lucide-svelte';
	import { goto } from '$app/navigation';

	function handleNewTopic() {
		goto('/forum/new');
	}
	function handleCategories() {
		goto('/forum/categories');
	}
</script>

<!-- Header Section -->
<div class="bg-gradient-to-r from-green-600 to-green-700 px-4 py-16">
	<div class="container mx-auto">
		<h1 class="mb-4 text-center text-3xl font-bold text-white md:text-4xl">Community Forum</h1>
		<p class="mx-auto mb-6 max-w-2xl text-center text-green-100">
			Connect with trainers worldwide. Share strategies, discuss updates, and build friendships.
		</p>

		<div class="flex justify-center">
			<button
				class="flex items-center rounded-lg bg-white px-6 py-3 font-semibold text-green-700 hover:bg-green-50"
				on:click={handleNewTopic}
			>
				<PlusCircle size={18} class="mr-2" />
				New Topic
			</button>
			<button
				class="flex items-center rounded-lg bg-white px-6 py-3 font-semibold text-green-700 hover:bg-green-50"
				on:click={handleCategories}
			>
				<Settings size={18} class="mr-2" />
				Manage Categories
			</button>
		</div>
	</div>
</div>

<!-- Main Content -->
<div class="container mx-auto px-4 py-12">
	<div class="grid grid-cols-1 gap-8 lg:grid-cols-4">
		<!-- Main Topics Column -->
		<div class="lg:col-span-3">
			<ForumControls />
			<ForumTopics />

			<!-- Pagination -->
			<nav class="mt-8 flex items-center justify-center">
				<button
					class="flex h-10 w-10 items-center justify-center rounded-l-lg border border-gray-300 text-gray-600 hover:bg-gray-100 disabled:opacity-50"
					disabled
				>
					Prev
				</button>
				<button
					class="flex h-10 w-10 items-center justify-center border-b border-t border-gray-300 bg-green-600 font-medium text-white"
				>
					1
				</button>
				<button
					class="flex h-10 w-10 items-center justify-center border-b border-t border-gray-300 text-gray-600 hover:bg-gray-100"
				>
					2
				</button>
				<button
					class="flex h-10 w-10 items-center justify-center border-b border-t border-gray-300 text-gray-600 hover:bg-gray-100"
				>
					3
				</button>
				<button
					class="flex h-10 w-10 items-center justify-center rounded-r-lg border border-gray-300 text-gray-600 hover:bg-gray-100"
				>
					Next
				</button>
			</nav>
		</div>

		<!-- Sidebar Column -->
		<ForumSidebar />
	</div>
</div>
