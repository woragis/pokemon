type Common = {
	name: string;
	url: string;
};
export interface Pokemon extends Common {}

export type Sprites = {
	front_default: string;
	front_shiny: string;
};

export type TypeName =
	| 'normal' // 1
	| 'fire' // 2
	| 'water' // 3
	| 'grass' // 4
	| 'steel' // 5
	| 'psychic' // 6
	| 'ghost' // 7
	| 'fairy' // 8
	| 'ground' // 9
	| 'rock' // 10
	| 'electric' // 11
	| 'fighting' // 12
	| 'flying' // 13
	| 'dark' // 14
	| 'dragon' // 15
	| 'ice' // 16
	| 'poison' // 17
	| 'bug' // 18
	| 'steelar' // 19
	| 'unknown' // 20
	| 'shadow'; // 21

export interface PokemonType {
	type: {
		name: TypeName;
		url: string;
	};
}

type Ability = {
	ability: Common;
	is_hidden: boolean;
	slot: 1 | 2 | 3;
};
type Stat = {
	base_stat: number;
	effort: number;
	stat: {
		name: stats;
		url: string;
	};
};
enum stats {
	hp = 'hp',
	attack = 'attack',
	defense = 'defense',
	special_attack = 'special-attack',
	special_defense = 'special-defense',
	speed = 'speed'
}
type Move = {
	move: {
		name: string;
		url: string;
	};
	version_group_details: {
		level_learned_at: number;
		move_learn_method: {
			name: string;
			url: string;
		};
		version_group: {
			name: string;
			url: string;
		};
	}[];
};
type Form = Common;
type Cry = {
	latest: string;
};
export interface PokemonData {
	name: string;
	id: number;
	sprites: Sprites;
	types: PokemonType[];

	weight: number;
	height: number;
	base_experience: number;
	abilities: Ability[];
	stats: Stat[];
	moves: Move[];
	forms: Form[];
	cries: Cry;
}

export type PokemonsInitialState = {
	pokemons: {
		list: Pokemon[];
		data: PokemonData[];
		search: PokemonData[];
		compare: PokemonData[];
	};
	types: {
		list: Type[];
		data: TypeData[];
	};
};

export interface PokemonsRequest {
	results: Pokemon[];
}

export interface Type extends Common {}

export type TypesSprites = {
	'generation-viii': {
		'sword-shield': {
			name_icon: string;
		};
	};
};

type apiObject = Common;

type DamageRelations = {
	double_damage_from: apiObject[];
	double_damage_to: apiObject[];
	half_damage_from: apiObject[];
	half_damage_to: apiObject[];
	no_damage_from: apiObject[];
	no_damage_to: apiObject[];
};

export interface TypeData {
	name: string;
	id: number;
	sprites: TypesSprites;
	damage_relations: DamageRelations;
}

export interface TypesRequest {
	results: Type[];
}

export type TypeEffectiveness = {
	weakTo: string[];
	strongAgainst: string[];
	resistantTo: string[];
	immuneTo: string[];
};
