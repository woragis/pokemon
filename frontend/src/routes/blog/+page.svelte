<script lang="ts">
	import { useBlogPostsQuery } from '$lib/api/blog';
	import { Clock, PenSquare, User } from 'lucide-svelte';

	const blogQuery = useBlogPostsQuery();
</script>

{#if $blogQuery.isLoading}
	<p class="center-text">Loading blog posts...</p>
{:else if $blogQuery.isError && $blogQuery.error}
	<p class="center-text error">Error: {$blogQuery.error.message}</p>
{:else}
	<section class="blog-hero">
		<h1>Pokémon Blog</h1>
		<p class="blog-subtitle">
			Discover in-depth articles, strategies, and stories from our community of trainers.
		</p>
		<a href="/blog/new" class="blog-btn">
			<PenSquare class="blog-icon" />
			Write New Post
		</a>
	</section>

	<section class="blog-posts">
		{#each $blogQuery.data as post}
			<div class="blog-post-card">
				<div class="blog-post-content">
					<h2>{post.title}</h2>
					<p>{post.content.slice(0, 150)}...</p>
					<div class="blog-post-meta">
						<User class="blog-icon" />
						<Clock class="blog-icon" />
					</div>
					<a href={`/blog/${post.id}`} class="blog-read-more">Read More →</a>
				</div>
			</div>
		{/each}
	</section>
{/if}

<style>
	.center-text {
		text-align: center;
		margin-top: 2rem;
	}

	.error {
		color: red;
	}

	.blog-hero {
		background: linear-gradient(to right, #2563eb, #1d4ed8);
		color: white;
		padding: 4rem 1rem;
		text-align: center;
	}

	.blog-hero h1 {
		font-size: 2.5rem;
		margin-bottom: 1rem;
	}

	.blog-subtitle {
		max-width: 600px;
		margin: 0 auto 2rem;
		font-size: 1.1rem;
		color: #cbd5e1;
	}

	.blog-btn {
		display: inline-flex;
		align-items: center;
		padding: 0.75rem 1.5rem;
		background-color: white;
		color: #2563eb;
		font-weight: bold;
		text-decoration: none;
		border-radius: 0.5rem;
		transition: background 0.2s ease-in-out;
	}

	.blog-btn:hover {
		background-color: #f1f5f9;
	}

	.blog-btn .blog-icon {
		margin-right: 0.5rem;
		width: 1rem;
		height: 1rem;
	}

	.blog-posts {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
		gap: 2rem;
		padding: 3rem 1rem;
		max-width: 1200px;
		margin: 0 auto;
	}

	.blog-post-card {
		background-color: white;
		border-radius: 0.75rem;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
		overflow: hidden;
	}

	.blog-post-content {
		padding: 1.5rem;
	}

	.blog-post-content h2 {
		font-size: 1.25rem;
		margin-bottom: 0.75rem;
		color: #111827;
	}

	.blog-post-content p {
		color: #4b5563;
		margin-bottom: 1rem;
	}

	.blog-post-meta {
		display: flex;
		align-items: center;
		color: #6b7280;
		font-size: 0.875rem;
		margin-bottom: 1rem;
	}

	.blog-post-meta .blog-icon {
		width: 1rem;
		height: 1rem;
		margin-right: 0.5rem;
	}

	.blog-read-more {
		color: #2563eb;
		font-weight: 500;
		text-decoration: none;
	}

	.blog-read-more:hover {
		text-decoration: underline;
	}
</style>
