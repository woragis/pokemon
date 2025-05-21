import { GAMES_URL } from '.';

export interface FetchGameProps {
	id: string;
}

export async function fetchGames() {
	try {
		const res = await fetch(`${GAMES_URL}/`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching games');
	}
}

export async function fetchGameById({ id }: FetchGameProps) {
	try {
		const res = await fetch(`${GAMES_URL}/${id}`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching game');
	}
}
