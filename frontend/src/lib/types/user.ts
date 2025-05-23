import type { PokemonGame } from './games';

export interface User {
	id: string;
	avatar: string;
	username: string;
	email: string;
	name: string;
	favorite_pokemon: string;
	favorite_game_id: string;
	favorite_game: PokemonGame;
	role: string;
	dob: string;
	created_at: string;
	updated_at: string;
	deleted_at: string;
}
