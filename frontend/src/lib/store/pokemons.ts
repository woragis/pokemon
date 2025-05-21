import type { PokemonData } from '$lib/types/pokeapi';
import { writable } from 'svelte/store';

export const pokemons = writable<PokemonData[]>([]);
export const total = writable<number>(0);
