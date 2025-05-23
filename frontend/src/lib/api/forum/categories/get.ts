import { FORUM_CATEGORIES_URL } from '.';

export interface FetchForumCategoryProps {
	id: string;
}

export async function fetchForumCategories() {
	try {
		const res = await fetch(`${FORUM_CATEGORIES_URL}/`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error fetching forum categories');
	}
}

export async function fetchForumCategoryById({ id }: FetchForumCategoryProps) {
	try {
		const res = await fetch(`${FORUM_CATEGORIES_URL}/${id}`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error fetching forum category');
	}
}
