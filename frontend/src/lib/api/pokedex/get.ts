import { POKEDEX_URL, type PokedexResponse } from '.';
import { getHeaders } from '..';

export async function fetchPokedex() {
	try {
		const res = await fetch(`${POKEDEX_URL}/`, {
			headers: getHeaders()
		});
		if (!res.ok) throw new Error();
		return (await res.json()) as PokedexResponse;
	} catch (e: any) {
		throw new Error('Error fetching pokedex');
	}
}
