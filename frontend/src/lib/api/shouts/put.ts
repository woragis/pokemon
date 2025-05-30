import type { Shout } from '$lib/types/shouts';
import { SHOUT_URL } from '.';
import { getHeaders } from '..';

export interface PutShoutProps {
	id: string;
	shout: { content: string };
}

export async function putShout({ id, shout }: PutShoutProps): Promise<Shout> {
	const headers = getHeaders();
	try {
		const res = await fetch(`${SHOUT_URL}/${id}`, {
			method: 'PUT',
			headers,
			body: JSON.stringify(shout)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error updating shout');
	}
}
