<script lang="ts">
	type Stat = {
		value: string;
		label: string;
	};

	const stats: Stat[] = [
		{ value: '890+', label: 'Pokémon Species' },
		{ value: '18', label: 'Pokémon Types' },
		{ value: '8', label: 'Generations' },
		{ value: '50K+', label: 'Active Trainers' }
	];
</script>

<section class="bg-gray-50 py-14">
	<div class="container mx-auto px-4">
		<div class="grid grid-cols-2 gap-8 md:grid-cols-4">
			{#each stats as stat}
				<div class="text-center">
					<div class="mb-2 text-4xl font-bold text-red-600 md:text-5xl">{stat.value}</div>
					<div class="text-gray-600">{stat.label}</div>
				</div>
			{/each}
		</div>
	</div>
</section>
