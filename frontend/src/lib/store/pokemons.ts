import type { Pagination } from '$lib/api/pokedex/get';
import type { PokemonData } from '$lib/types/pokeapi';
import { derived, writable } from 'svelte/store';

export const pokemons = writable<PokemonData[]>([]);
export const total = 1025;
export const pagination = writable<Pagination>({ offset: 0, limit: 20 });
export const pageNumbers = derived([pagination], ([$pagination]) => {
	const totalPages = Math.ceil(total / $pagination.limit);
	const currentPage = $pagination.offset / $pagination.limit + 1;

	let start = Math.max(1, currentPage - 2);
	let end = Math.min(totalPages, start + 4);
	start = Math.max(1, end - 4); // shift range back if near the end

	return Array.from({ length: end - start + 1 }, (_, i) => start + i);
});
