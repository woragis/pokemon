import { API_URL } from '..';
import type { DefaultResponse } from '$lib/types';
import type { PokedexPokemon } from '$lib/types/pokemon';
import { fetchPokedex } from './get';
export const POKEDEX_URL = `${API_URL}/pokedex`;

export interface PokedexResponse extends DefaultResponse {
	pokemons: PokedexPokemon[];
}

export { fetchPokedex };
