import type { DefaultResponse, PokedexPokemon } from '@/lib/types';
import { API_URL, getHeaders } from '.';

const POKEDEX_URL = `${API_URL}/pokedex`;

export async function fetchPokedex() {
	try {
		const res = await fetch(`${POKEDEX_URL}/`);
		if (!res.ok) throw new Error();
		return (await res.json()) as PokedexResponse;
	} catch (e: any) {
		throw new Error('Error fetching pokedex');
	}
}

interface PokedexResponse extends DefaultResponse {
	pokemons: PokedexPokemon[];
}
