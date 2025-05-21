import { getHeaders } from '..';
import { BLOG_URL } from '.';
import type { BlogPost } from '$lib/types/blog';

export interface PostBlogPostProps {
	blog: Omit<BlogPost, 'id'>;
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
