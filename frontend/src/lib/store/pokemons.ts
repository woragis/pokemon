import type { Pagination } from '$lib/api/pokedex/get';
import type { PokemonData } from '$lib/types/pokeapi';
import { writable } from 'svelte/store';

export const pokemons = writable<PokemonData[]>([]);
export const total = writable<number>(0);
export const pagination = writable<Pagination>({ offset: 0, limit: 20 });
