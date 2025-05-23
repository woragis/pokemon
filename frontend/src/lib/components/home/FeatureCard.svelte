<script lang="ts">
	import { ChevronRight } from 'lucide-svelte';
	import { onMount } from 'svelte';

	export let title;
	export let description;
	export let icon; // This should be a Svelte component or raw SVG passed as a slot or prop
	export let color;
	export let link;

	let element: HTMLElement;
	let visible = false;

	onMount(() => {
		const observer = new IntersectionObserver(
			([entry]) => {
				if (entry.isIntersecting) {
					visible = true;
					observer.disconnect();
				}
			},
			{ threshold: 0.1 }
		);

		if (element) observer.observe(element);
	});
</script>

<div
	class="feature-card"
	bind:this={element}
	class:animate={visible}
	style={`--card-color: ${color}`}
>
	<div class={`feature-card-header ${color}`}>
		<div class="feature-icon-wrapper">
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

	<div class="feature-card-body">
		<h3 class="feature-card-title">{title}</h3>
		<p class="feature-card-description">{description}</p>
	</div>

	<div class="feature-card-footer">
		<a href={link} class="feature-card-link">
			Learn more <ChevronRight size="16" class="chevron-icon" />
		</a>
	</div>
</div>

<style>
	@keyframes fadeInUp {
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.feature-card {
		display: flex;
		flex-direction: column;
		height: 100%;
		overflow: hidden;
		border-radius: 12px;
		background-color: white;
		box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
		transition:
			box-shadow 0.3s ease,
			transform 0.3s ease;
		opacity: 0;
		transform: translateY(20px);
	}

	.feature-card.animate {
		animation: fadeInUp 0.5s ease forwards;
	}

	.feature-card:hover {
		box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
		transform: translateY(-4px);
	}

	.feature-card-header {
		padding: 1.5rem;
		background-color: var(--card-color, #e2e8f0); /* fallback */
		transition: background-color 0.3s ease;
	}

	.feature-icon-wrapper {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		padding: 0.75rem;
		border-radius: 9999px;
		background-color: rgba(255, 255, 255, 0.2);
		backdrop-filter: blur(4px);
	}

	.feature-card-body {
		flex-grow: 1;
		padding: 1.5rem;
	}

	.feature-card-title {
		font-size: 1.25rem;
		font-weight: 600;
		color: #1f2937; /* gray-900 */
		margin-bottom: 0.75rem;
	}

	.feature-card-description {
		color: #4b5563; /* gray-600 */
		margin-bottom: 1rem;
	}

	.feature-card-footer {
		padding: 0 1.5rem 1.5rem 1.5rem;
	}

	.feature-card-link {
		display: inline-flex;
		align-items: center;
		font-size: 0.875rem;
		font-weight: 500;
		color: #2563eb; /* blue-600 */
		text-decoration: none;
		transition: color 0.2s;
	}

	.feature-card-link:hover {
		color: #1e40af; /* blue-800 */
	}

	.chevron-icon {
		margin-left: 0.25rem;
	}
</style>
