import { getHeaders } from '..';
import { GUIDE_URL } from '.';

export interface DeleteGuideProps {
	id: string;
}

export async function deleteGuide({ id }: DeleteGuideProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${GUIDE_URL}/${id}`, {
			method: 'DELETE',
			headers
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error deleting guide');
	}
}
