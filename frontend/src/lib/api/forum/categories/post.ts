import { getHeaders } from '../..';
import { FORUM_CATEGORIES_URL } from '.';
import type { ForumCategory } from '$lib/types/forum';

export interface PostForumCategoryProps {
	category: Omit<ForumCategory, 'id'>;
}

export async function postForumCategory({ category }: PostForumCategoryProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${FORUM_CATEGORIES_URL}/`, {
			method: 'POST',
			headers,
			body: JSON.stringify(category)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error creating forum category');
	}
}
