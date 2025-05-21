import { getHeaders } from '..';
import { BLOG_URL } from '.';

export interface DeleteBlogPostProps {
	id: string;
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
