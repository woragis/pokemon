import { API_URL, getHeaders } from '.';
import type { DefaultResponse } from '../types';
import type { PokedexPokemon } from '../types/pokemon';

const POKEDEX_URL = `${API_URL}/pokedex`;

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

interface PokedexResponse extends DefaultResponse {
	pokemons: PokedexPokemon[];
}
