import { getHeaders } from '..';
import { GAMES_URL } from '.';
import type { PokemonGame } from '$lib/types/games';

export interface PutGameProps {
	id: string;
	game: Omit<PokemonGame, 'id'>;
}

export async function putGame({ id, game }: PutGameProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${GAMES_URL}/${id}`, {
			method: 'PUT',
			headers,
			body: JSON.stringify(game)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error updating game');
	}
}
