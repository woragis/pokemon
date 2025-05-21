<script>
	import { ChevronRight } from 'lucide-svelte';

	export let title;
	export let description;
	export let icon; // This should be a Svelte component or raw SVG passed as a slot or prop
	export let color;
	export let link;
</script>

<div
	class="flex h-full flex-col overflow-hidden rounded-xl bg-white shadow-lg transition-shadow duration-300 hover:shadow-xl"
>
	<div class={`p-6 ${color}`}>
		<div
			class="inline-flex items-center justify-center rounded-full bg-white/20 p-3 backdrop-blur-sm"
		>
			{#if icon}
				{#await icon}
					<!-- fallback if icon is async -->
				{:then Icon}
					<Icon />
				{/await}
			{:else}
				{icon}
			{/if}
		</div>
	</div>

	<div class="flex-grow p-6">
		<h3 class="mb-3 text-xl font-semibold text-gray-900">{title}</h3>
		<p class="mb-4 text-gray-600">{description}</p>
	</div>

	<div class="px-6 pb-6">
		<a
			href={link}
			class="inline-flex items-center text-sm font-medium text-blue-600 transition-colors hover:text-blue-800"
		>
			Learn more <ChevronRight size="16" class="ml-1" />
		</a>
	</div>
</div>
