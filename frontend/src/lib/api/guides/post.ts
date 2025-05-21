import { getHeaders } from '..';
import { GUIDE_URL } from '.';
import type { Guide } from '$lib/types/guides';

export interface PostGuideProps {
	guide: Omit<Guide, 'id'>;
}

export async function postGuide({ guide }: PostGuideProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${GUIDE_URL}/`, {
			method: 'POST',
			headers,
			body: JSON.stringify(guide)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error posting guide');
	}
}
