import { API_URL, getHeaders } from '.';
import type { BlogPost } from '../types/blog';

const BLOG_URL = `${API_URL}/blog`;

export async function fetchBlogPosts() {
	try {
		const res = await fetch(`${BLOG_URL}/`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching blog posts');
	}
}

export async function fetchBlogPost({ id }: FetchBlogPostProps) {
	try {
		const res = await fetch(`${BLOG_URL}/${id}`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching blog post');
	}
}

export async function postBlogPost({ blog }: PostBlogPostProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${BLOG_URL}/`, {
			method: 'POST',
			headers,
			body: JSON.stringify(blog)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching blog post');
	}
}

export async function putBlogPost({ id, blog }: PutBlogPostProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${BLOG_URL}/${id}`, {
			method: 'PUT',
			headers,
			body: JSON.stringify(blog)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching blog post');
	}
}

export async function deleteBlogPost({ id }: DeleteBlogPostProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${BLOG_URL}/${id}`, {
			method: 'DELETE',
			headers
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching blog post');
	}
}

interface FetchBlogPostProps {
	id: string;
}
interface PostBlogPostProps {
	blog: Omit<BlogPost, 'id'>;
}
interface PutBlogPostProps {
	id: string;
	blog: Omit<BlogPost, 'id'>;
}
interface DeleteBlogPostProps {
	id: string;
}
