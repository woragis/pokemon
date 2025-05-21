import { BLOG_URL } from '.';

export interface FetchBlogPostProps {
	id: string;
}

export async function fetchBlogPosts() {
	try {
		const res = await fetch(`${BLOG_URL}/`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching blog posts');
	}
}

export async function fetchBlogPostById({ id }: FetchBlogPostProps) {
	try {
		const res = await fetch(`${BLOG_URL}/${id}`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching blog post');
	}
}
