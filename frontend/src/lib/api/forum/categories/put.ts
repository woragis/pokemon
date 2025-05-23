import { getHeaders } from '../..';
import { FORUM_CATEGORIES_URL } from '.';
import type { ForumCategory } from '$lib/types/forum';

export interface PutForumCategoryProps {
	id: string;
	category: Omit<ForumCategory, 'id'>;
}

export async function putForumCategory({ id, category }: PutForumCategoryProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${FORUM_CATEGORIES_URL}/${id}`, {
			method: 'PUT',
			headers,
			body: JSON.stringify(category)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error updating forum category');
	}
}
