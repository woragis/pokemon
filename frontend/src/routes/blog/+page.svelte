<script lang="ts">
	import { onMount } from 'svelte';
	import type { BlogPost } from '$lib/types/blog';
	import { fetchBlogPosts } from '../../lib/api/blog';
	import { Clock, PenSquare, User } from 'lucide-svelte';

	let posts: BlogPost[] = [];
	let error: string | null = null;
	let loading = false;

	async function handleFetch() {
		loading = true;
		try {
			posts = await fetchBlogPosts();
		} catch (e: any) {
			error = e.message;
		} finally {
			loading = false;
		}
	}

	onMount(() => {
		handleFetch();
	});
</script>

<h1>Pokémon Blog</h1>

{#if loading}
	<p>Loading blog posts...</p>
{:else if error}
	<p>Error: {error}</p>
{:else}
	<ul>
		{#each posts as post}
			<li>
				<a href={`/blog/${post.id}`}>
					<h2>{post.title}</h2>
					<small>{post.date}</small>
					<p>{post.excerpt}</p>
				</a>
			</li>
		{/each}
	</ul>
{/if}

<div class="bg-gradient-to-r from-blue-600 to-blue-700 px-4 py-16">
	<div class="container mx-auto">
		<h1 class="mb-4 text-center text-3xl font-bold text-white md:text-4xl">Pokémon Blog</h1>
		<p class="mx-auto mb-8 max-w-2xl text-center text-blue-100">
			Discover in-depth articles, strategies, and stories from our community of trainers.
		</p>
		<div class="flex justify-center">
			<a
				href="/blog/new"
				class="flex items-center rounded-lg bg-white px-6 py-3 font-semibold text-blue-600 transition-colors hover:bg-blue-50"
			>
				<PenSquare class="mr-1 h-4 w-4" />
				Write New Post
			</a>
		</div>
	</div>
</div>

<div class="container mx-auto px-4 py-12">
	<div class="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
		{#each posts as post}
			<div class="overflow-hidden rounded-lg bg-white shadow-md">
				<div class="p-6">
					<h2 class="mb-2 text-xl font-bold text-gray-900">{post.title}</h2>
					<p class="mb-4 text-gray-600">{post.content.slice(0, 150)}...</p>
					<div class="mb-4 flex items-center text-sm text-gray-500">
						<User class="mr-1 h-4 w-4" />
						<!-- <span class="mr-4">{post.author.username}</span> -->
						<Clock class="mr-1 h-4 w-4" />
						<!-- <span>{format(new Date(post.created_at), 'MMM d, yyyy')}</span> -->
					</div>
					<a
						href={`/blog/${post.id}`}
						class="font-medium text-blue-600 transition-colors hover:text-blue-800">Read More →</a
					>
				</div>
			</div>
		{/each}
	</div>
</div>
