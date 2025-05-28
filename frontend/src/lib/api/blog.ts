import { API_URL } from '.';
import { getHeaders } from '.';
import { createMutation, createQuery, useQueryClient } from '@tanstack/svelte-query';
import type { BlogPost } from '$lib/types/blog';

const BLOG_URL = `${API_URL}/blog`;

// Tanstack Hooks

export function useBlogPostsQuery() {
	return createQuery({
		queryKey: ['blogPosts'],
		queryFn: fetchBlogPosts
	});
}

export function useBlogPostQuery(id: string) {
	return createQuery({
		queryKey: ['blogPost', id],
		queryFn: () => fetchBlogPostById({ id }),
		enabled: !!id // only fetch if id exists
	});
}

export function useCreateBlogPostMutation() {
	const queryClient = useQueryClient();
	return createMutation({
		mutationFn: postBlogPost,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ['blogPosts'] });
		}
	});
}

export function useUpdateBlogPostMutation() {
	const queryClient = useQueryClient();
	return createMutation({
		mutationFn: putBlogPost,
		onSuccess: (_data, { id }) => {
			queryClient.invalidateQueries({ queryKey: ['blogPost', id, 'blogPosts'] });
		}
	});
}

export function useDeleteBlogPostMutation() {
	const queryClient = useQueryClient();
	return createMutation({
		mutationFn: deleteBlogPost,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ['blogPosts'] });
		}
	});
}

// Functions
async function fetchBlogPosts() {
	const res = await fetch(`${BLOG_URL}/`);
	if (!res.ok) throw new Error('Error fetching blog posts');
	return await res.json();
}

interface FetchBlogPostProps {
	id: string;
}

async function fetchBlogPostById({ id }: FetchBlogPostProps): Promise<BlogPost> {
	const res = await fetch(`${BLOG_URL}/${id}`);
	if (!res.ok) throw new Error('Error fetching blog post');
	return await res.json();
}

interface PostBlogPostProps {
	blog: Omit<BlogPost, 'id'>;
}
async function postBlogPost({ blog }: PostBlogPostProps) {
	const headers = getHeaders();
	const res = await fetch(`${BLOG_URL}/`, {
		method: 'POST',
		headers,
		body: JSON.stringify(blog)
	});
	if (!res.ok) throw new Error('Error fetching blog post');
	return await res.json();
}

interface PutBlogPostProps {
	id: string;
	blog: Omit<BlogPost, 'id'>;
}

async function putBlogPost({ id, blog }: PutBlogPostProps) {
	const headers = getHeaders();
	const res = await fetch(`${BLOG_URL}/${id}`, {
		method: 'PUT',
		headers,
		body: JSON.stringify(blog)
	});
	if (!res.ok) throw new Error('Error fetching blog post');
	return await res.json();
}

interface DeleteBlogPostProps {
	id: string;
}

async function deleteBlogPost({ id }: DeleteBlogPostProps) {
	const headers = getHeaders();
	const res = await fetch(`${BLOG_URL}/${id}`, {
		method: 'DELETE',
		headers
	});
	if (!res.ok) throw new Error('Error fetching blog post');
	return await res.json();
}
