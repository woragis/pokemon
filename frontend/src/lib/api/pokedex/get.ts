import { pokemons, total } from '$lib/store/pokemons';
import type { PokemonData } from '$lib/types/pokeapi';
import { POKEMONS_POKEAPI } from '.';

export interface Pagination {
	offset: number;
	limit: number;
}
type Common = {
	name: string;
	url: string;
};
type Pokemon = Common;
interface AllPokemonsResponse {
	count: number;
	results: Pokemon[];
}

export async function fetchAllPokemons({ offset = 0, limit = 100 }: Pagination) {
	try {
		const res = await fetch(`${POKEMONS_POKEAPI}?offset=${offset}&limit=${limit}`);
		if (!res.ok) throw new Error('Failed to fetch Pokémon list');
		const response: AllPokemonsResponse = await res.json();
		total.set(response.count);

		const details = await Promise.all(
			response.results.map(async (pokemon) => {
				const res = await fetch(pokemon.url);
				if (!res.ok) throw new Error('Failed to fetch Pokémon data');
				return (await res.json()) as PokemonData;
			})
		);

		pokemons.set(details); // use `.set()` instead of push

		return response;
	} catch (e: any) {
		throw new Error('Error fetching Pokédex');
	}
}
