<script lang="ts">
	import { onMount } from 'svelte';
	import { ArrowLeft, Save } from 'lucide-svelte';
	import { fetchBlogPostById } from '$lib/api/blog';

	export let id: string | undefined = undefined;

	let error = '';
	let post = {
		title: '',
		content: '',
		published: false
	};

	let loading = false;
	const isEditing = !id;

	async function handleFetch() {
		try {
			if (id) {
				post = await fetchBlogPostById({ id });
			}
		} catch (err) {
			error = '';
		}
	}

	const handleEdit = async () => {
		loading = true;
		try {
			// await putBlogPost({id, blog: post})
		} catch (error) {
			console.error('Error editing post:', error);
		} finally {
			loading = false;
		}
	};

	onMount(handleFetch);
</script>

<div class="min-h-screen bg-gray-50 pt-16">
	<div class="container mx-auto px-4 py-8">
		<div class="mx-auto max-w-4xl">
			<div class="mb-6">
				<a href="/blog" class="inline-flex items-center text-gray-600 hover:text-gray-900">
					<ArrowLeft class="mr-2 h-4 w-4" />
					Back to Blog
				</a>
			</div>

			{#if isEditing}
				<div class="rounded-lg bg-white p-6 shadow-md">
					<h1 class="mb-6 text-2xl font-bold text-gray-900">Create New Post</h1>
					<form on:submit|preventDefault={handleEdit}>
						<div class="mb-4">
							<label class="mb-1 block text-sm font-medium text-gray-700">Title</label>
							<input
								type="text"
								bind:value={post.title}
								class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
								required
							/>
						</div>

						<div class="mb-4">
							<label class="mb-1 block text-sm font-medium text-gray-700"
								>Content (Markdown supported)</label
							>
							<textarea
								bind:value={post.content}
								rows="15"
								class="w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
								required
							/>
						</div>

						<div class="mb-6">
							<label class="flex items-center">
								<input
									type="checkbox"
									bind:checked={post.published}
									class="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
								/>
								<span class="ml-2 text-sm text-gray-600">Publish immediately</span>
							</label>
						</div>

						<div class="flex justify-end">
							<button
								type="submit"
								class="flex items-center rounded-md bg-blue-600 px-6 py-2 text-white transition-colors hover:bg-blue-700"
								disabled={loading}
							>
								<Save class="mr-2 h-4 w-4" />
								{loading ? 'Saving...' : 'Save Post'}
							</button>
						</div>
					</form>
				</div>
			{:else}
				<div class="rounded-lg bg-white p-6 shadow-md">
					<h1 class="mb-4 text-3xl font-bold text-gray-900">{post.title}</h1>
					<!-- <div class="prose max-w-none" innerHTML={marked(post.content)} /> -->
				</div>
			{/if}
		</div>
	</div>
</div>
