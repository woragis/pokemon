import { getHeaders } from '../..';
import { FORUM_CATEGORIES_URL } from '.';

export interface DeleteForumCategoryProps {
	id: string;
}

export async function deleteForumCategory({ id }: DeleteForumCategoryProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${FORUM_CATEGORIES_URL}/${id}`, {
			method: 'DELETE',
			headers
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error deleting forum category');
	}
}
