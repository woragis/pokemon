import { getHeaders } from '..';
import { GAMES_URL } from '.';

export interface DeleteGameProps {
	id: string;
}

export async function deleteGame({ id }: DeleteGameProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${GAMES_URL}/${id}`, {
			method: 'DELETE',
			headers
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error deleting game');
	}
}
