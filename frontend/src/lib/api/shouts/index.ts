import { fetchShouts, fetchShoutById } from './get';
import { postShout } from './post';
import { putShout } from './put';
import { deleteShout } from './delete';
import { API_URL } from '..';

export const SHOUT_URL = `${API_URL}/shouts`;

export { fetchShouts, fetchShoutById, postShout, putShout, deleteShout };
