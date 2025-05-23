import { getHeaders } from '..';
import { SHOUT_URL } from '.';

export interface DeleteShoutProps {
	id: string;
}

export async function deleteShout({ id }: DeleteShoutProps): Promise<void> {
	const headers = getHeaders();
	try {
		const res = await fetch(`${SHOUT_URL}/${id}`, {
			method: 'DELETE',
			headers
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error deleting shout');
	}
}
