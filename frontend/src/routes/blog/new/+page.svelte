<script lang="ts">
	import { PenSquare, ArrowLeft } from 'lucide-svelte';
	import { goto } from '$app/navigation';
	import { useCreateBlogPostMutation } from '$lib/api/blog';

	const mutation = useCreateBlogPostMutation();
	let error = '';
	let title = '';
	let content = '';
	let published = false;
	let loading = false;

	const handleSubmit = async () => {
		if ($mutation.isPending || loading) return;
		loading = true;
		$mutation.mutate(
			{
				blog: {
					title,
					content,
					date: new Date().getUTCMilliseconds(),
					excerpt: ''
				}
			},
			{
				onSuccess: () => {
					goto('/blog');
				},
				onError: (err) => {
					error = `Failed to create blog post: ${err.message || err}`;
				}
			}
		);
		loading = false;
	};
</script>

<section class="blog-hero">
	{#if error}
		<p class="error">
			{error}
		</p>
	{/if}
	<h1>Create New Post</h1>
	<p class="blog-subtitle">Write and publish your Pokémon blog post to share with the community.</p>
	<a href="/blog" class="blog-btn">
		<ArrowLeft class="blog-icon" />
		Back to Blog
	</a>
</section>

<section class="blog-posts">
	<form class="form-card" on:submit|preventDefault={handleSubmit}>
		<div class="form-group">
			<label for="title">Title</label>
			<input
				id="title"
				type="text"
				bind:value={title}
				required
				placeholder="Enter your post title"
			/>
		</div>

		<div class="form-group">
			<label for="content">Content (Markdown supported)</label>
			<textarea
				id="content"
				rows="10"
				bind:value={content}
				required
				placeholder="Write your post content here..."
			></textarea>
		</div>

		<div class="form-checkbox">
			<label>
				<input type="checkbox" bind:checked={published} />
				<span>Publish immediately</span>
			</label>
		</div>

		<div class="form-actions">
			<button type="submit" disabled={loading} class="blog-btn">
				<PenSquare class="blog-icon" />
				{loading ? 'Saving...' : 'Save Post'}
			</button>
		</div>
	</form>
</section>

<style>
	/* Reuse existing styles plus some additions */

	.form-card {
		background-color: white;
		border-radius: 0.75rem;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
		max-width: 700px;
		margin: 2rem auto 4rem;
		padding: 2rem;
	}

	.form-group {
		margin-bottom: 1.5rem;
		display: flex;
		flex-direction: column;
	}

	label {
		font-weight: 600;
		color: #111827;
		margin-bottom: 0.5rem;
	}

	input[type='text'],
	textarea {
		padding: 0.5rem 0.75rem;
		border: 1px solid #d1d5db;
		border-radius: 0.375rem;
		font-size: 1rem;
		color: #111827;
		resize: vertical;
	}

	input[type='text']:focus,
	textarea:focus {
		outline: none;
		border-color: #2563eb;
		box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.3);
	}

	.form-checkbox {
		display: flex;
		align-items: center;
		margin-bottom: 1.5rem;
		color: #374151;
		font-weight: 500;
	}

	.form-checkbox input[type='checkbox'] {
		margin-right: 0.5rem;
		width: 1.125rem;
		height: 1.125rem;
	}

	.form-actions {
		text-align: right;
	}

	.form-actions button.blog-btn {
		padding: 0.5rem 1.5rem;
		font-size: 1rem;
	}

	.form-actions button.blog-btn:disabled {
		opacity: 0.6;
		cursor: not-allowed;
	}
</style>
