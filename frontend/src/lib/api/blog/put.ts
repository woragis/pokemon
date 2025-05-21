import type { BlogPost } from '$lib/types/blog';
import { BLOG_URL } from '.';
import { getHeaders } from '..';

export interface PutBlogPostProps {
	id: string;
	blog: Omit<BlogPost, 'id'>;
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
