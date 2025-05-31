<script lang="ts">
	import { ArrowLeft, Save } from 'lucide-svelte';
	import { useBlogPostQuery } from '$lib/api/blog';
	import type { BlogPost } from '$lib/types/blog';
	import { onMount } from 'svelte';

	export let id: string | undefined = undefined;
	const isEditing = !id;

	const postQuery = id ? useBlogPostQuery(id) : undefined;

	let localPost: BlogPost = {
		id: '',
		title: '',
		content: '',
		date: '',
		excerpt: ''
	};

	let loading = false;

	onMount(() => {
		if ($postQuery && $postQuery.data) {
			// Clone the data into the local editable object
			localPost = { ...$postQuery.data };
		}
	});

	const handleEdit = async () => {
		loading = true;
		try {
			// Send `localPost` to backend
			console.log(localPost);
			// await putBlogPost({ id, blog: localPost });
		} catch (error) {
			console.error('Error editing post:', error);
		} finally {
			loading = false;
		}
	};
</script>

{#if $postQuery && $postQuery.isLoading}
	<p class="center-text">Loading blog post...</p>
{:else if $postQuery && $postQuery.isError && $postQuery.error}
	<p class="center-text error">Error: {$postQuery.error.message}</p>
{:else}
	<div class="page">
		<div class="container">
			<div class="back-link">
				<a href="/blog">
					<ArrowLeft class="icon" />
					Back to Blog
				</a>
			</div>

			{#if isEditing}
				<div class="form-card">
					<h1 class="form-title">Create New Post</h1>
					<form on:submit|preventDefault={handleEdit}>
						<div class="form-group">
							<label>Title</label>
							<input type="text" bind:value={localPost.title} required />
						</div>

						<div class="form-group">
							<label>Content (Markdown supported)</label>
							<textarea rows="15" bind:value={localPost.content} required />
						</div>

						<div class="form-checkbox">
							<!-- <label>
								<input type="checkbox" bind:checked={localPost.excerpt} />
								<span>Publish immediately</span>
							</label> -->
						</div>

						<div class="form-actions">
							<button type="submit" disabled={loading}>
								<Save class="icon" />
								{loading ? 'Saving...' : 'Save Post'}
							</button>
						</div>
					</form>
				</div>
			{:else if $postQuery && $postQuery.data}
				<div class="form-card">
					<h1 class="form-title">{$postQuery.data.title}</h1>
					<!-- <div class="prose" innerHTML={marked($postQuery.data.content)} /> -->
				</div>
			{/if}
		</div>
	</div>
{/if}

<style>
	.page {
		min-height: 100vh;
		background-color: #f9fafb;
		padding-top: 4rem;
	}

	.container {
		max-width: 960px;
		margin: 0 auto;
		padding: 2rem;
	}

	.back-link {
		margin-bottom: 1.5rem;
	}

	.back-link a {
		display: inline-flex;
		align-items: center;
		color: #4b5563;
		text-decoration: none;
		font-weight: 500;
		transition: color 0.2s;
	}

	.back-link a:hover {
		color: #1f2937;
	}

	.icon {
		margin-right: 0.5rem;
		width: 1rem;
		height: 1rem;
	}

	.form-card {
		background-color: white;
		border-radius: 0.5rem;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
		padding: 2rem;
	}

	.form-title {
		font-size: 1.5rem;
		font-weight: bold;
		color: #111827;
		margin-bottom: 1.5rem;
	}

	.form-group {
		margin-bottom: 1.5rem;
	}

	.form-group label {
		display: block;
		font-size: 0.875rem;
		font-weight: 500;
		margin-bottom: 0.5rem;
		color: #374151;
	}

	.form-group input,
	.form-group textarea {
		width: 100%;
		padding: 0.75rem;
		border: 1px solid #d1d5db;
		border-radius: 0.375rem;
		font-size: 1rem;
	}

	.form-group input:focus,
	.form-group textarea:focus {
		border-color: #3b82f6;
		outline: none;
		box-shadow: 0 0 0 1px #3b82f6;
	}

	.form-checkbox {
		margin-bottom: 2rem;
	}

	.form-checkbox label {
		display: flex;
		align-items: center;
		font-size: 0.875rem;
		color: #4b5563;
	}

	.form-checkbox input[type='checkbox'] {
		margin-right: 0.5rem;
	}

	.form-actions {
		display: flex;
		justify-content: flex-end;
	}

	.form-actions button {
		display: inline-flex;
		align-items: center;
		padding: 0.5rem 1.5rem;
		background-color: #2563eb;
		color: white;
		border: none;
		border-radius: 0.375rem;
		font-size: 1rem;
		cursor: pointer;
		transition: background-color 0.2s;
	}

	.form-actions button:hover {
		background-color: #1d4ed8;
	}

	.form-actions button:disabled {
		background-color: #93c5fd;
		cursor: not-allowed;
	}
</style>
