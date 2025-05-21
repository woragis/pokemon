import { getHeaders } from '..';
import { GAMES_URL } from '.';
import type { PokemonGame } from '$lib/types/games';

export interface PostGameProps {
	game: Omit<PokemonGame, 'id'>;
}

export async function postGame({ game }: PostGameProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${GAMES_URL}/`, {
			method: 'POST',
			headers,
			body: JSON.stringify(game)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error creating game');
	}
}
