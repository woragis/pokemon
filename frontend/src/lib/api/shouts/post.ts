import { getHeaders } from '..';
import { SHOUT_URL } from '.';
import type { Shout } from '$lib/types/shouts';

export interface PostShoutProps {
	shout: { content: string };
}

export async function postShout({ shout }: PostShoutProps): Promise<Shout> {
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
