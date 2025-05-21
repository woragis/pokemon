import { fetchGames, fetchGameById } from './get';
import { postGame } from './post';
import { putGame } from './put';
import { deleteGame } from './delete';
import { API_URL } from '..';

export const GAMES_URL = `${API_URL}/games`;

export { fetchGames, fetchGameById, postGame, putGame, deleteGame };
