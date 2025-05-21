import type { Guide } from '$lib/types/guides';
import { GUIDE_URL } from '.';
import { getHeaders } from '..';

export interface PutGuideProps {
	id: string;
	guide: Omit<Guide, 'id'>;
}

export async function putGuide({ id, guide }: PutGuideProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${GUIDE_URL}/${id}`, {
			method: 'PUT',
			headers,
			body: JSON.stringify(guide)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error updating guide');
	}
}
