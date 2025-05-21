import { getHeaders } from '..';
import { SHOUT_URL } from '.';
import type { Shout } from '$lib/types/shouts';

export interface PostShoutProps {
	shout: Omit<Shout, 'id' | 'created_at' | 'updated_at'>;
}

export async function postShout({ shout }: PostShoutProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${SHOUT_URL}/`, {
			method: 'POST',
			headers,
			body: JSON.stringify(shout)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error posting shout');
	}
}
