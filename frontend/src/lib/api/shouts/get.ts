import type { Shout } from '$lib/types/shouts';
import { SHOUT_URL } from '.';

export interface FetchShoutProps {
	id: string;
}

export async function fetchShouts(): Promise<Shout[]> {
	try {
		const res = await fetch(`${SHOUT_URL}/`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching shouts');
	}
}

export async function fetchShoutById({ id }: FetchShoutProps): Promise<Shout> {
	try {
		const res = await fetch(`${SHOUT_URL}/${id}`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching shout');
	}
}
