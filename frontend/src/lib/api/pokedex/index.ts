import { API_URL } from '..';
import type { DefaultResponse } from '$lib/types';
import type { PokedexPokemon } from '$lib/types/pokemon';
import { fetchAllPokemons } from './get';
export const POKEDEX_URL = `${API_URL}/pokedex`;
export const POKEAPI = 'https://pokeapi.co/api/v2';
export const POKEMONS_POKEAPI = `${POKEAPI}/pokemon`;

export interface PokedexResponse extends DefaultResponse {
	pokemons: PokedexPokemon[];
}

export { fetchAllPokemons };
