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
export const typeColors: Record<string, string> = {
	normal: 'bg-gray-300',
	fire: 'bg-red-500',
	water: 'bg-blue-500',
	electric: 'bg-yellow-400',
	grass: 'bg-green-500',
	ice: 'bg-blue-200',
	fighting: 'bg-red-700',
	poison: 'bg-purple-500',
	ground: 'bg-yellow-600',
	flying: 'bg-indigo-300',
	psychic: 'bg-pink-500',
	bug: 'bg-green-400',
	rock: 'bg-yellow-700',
	ghost: 'bg-purple-700',
	dragon: 'bg-indigo-600',
	dark: 'bg-gray-700',
	steel: 'bg-gray-400',
	fairy: 'bg-pink-300'
};
