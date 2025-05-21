<script lang="ts">
	import { pageNumbers, pagination, total } from '$lib/store/pokemons';
	import { derived } from 'svelte/store';

	const currentPage = derived(pagination, ($p) => Math.floor($p.offset / $p.limit) + 1);
	const totalPages = derived(pagination, ($p) => Math.ceil(total / $p.limit));

	function goToPage(page: number) {
		pagination.update((p) => ({ ...p, offset: (page - 1) * p.limit }));
	}

	function prevPage() {
		pagination.update((p) => (p.offset === 0 ? p : { ...p, offset: p.offset - p.limit }));
	}

	function nextPage() {
		pagination.update((p) =>
			p.offset + p.limit >= total ? p : { ...p, offset: p.offset + p.limit }
		);
	}
</script>

<div class="mt-12 flex justify-center">
	<nav class="flex flex-wrap items-center gap-2">
		<!-- First Page -->
		<button
			on:click={() => goToPage(1)}
			class="rounded-full bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm ring-1 ring-gray-300 hover:bg-red-50 disabled:opacity-50"
			disabled={$currentPage === 1}
		>
			« First
		</button>

		<!-- Prev -->
		<button
			on:click={prevPage}
			class="rounded-full bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm ring-1 ring-gray-300 hover:bg-red-50 disabled:opacity-50"
			disabled={$currentPage === 1}
		>
			← Prev
		</button>

		<!-- Page Numbers -->
		{#each $pageNumbers as page}
			<button
				on:click={() => goToPage(page)}
				class={`rounded-full px-4 py-2 text-sm font-medium shadow ${
					page === $currentPage
						? 'bg-red-600 text-white'
						: 'bg-white text-gray-700 ring-1 ring-gray-300 hover:bg-red-50'
				}`}
			>
				{page}
			</button>
		{/each}

		<!-- Next -->
		<button
			on:click={nextPage}
			class="rounded-full bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm ring-1 ring-gray-300 hover:bg-red-50 disabled:opacity-50"
			disabled={$currentPage === $totalPages}
		>
			Next →
		</button>

		<!-- Last Page -->
		<button
			on:click={() => goToPage($totalPages)}
			class="rounded-full bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm ring-1 ring-gray-300 hover:bg-red-50 disabled:opacity-50"
			disabled={$currentPage === $totalPages}
		>
			Last »
		</button>
	</nav>
</div>
