import { API_URL } from '..';
import { getAllSnaps, getMySnaps, getSnap } from './get';
import { createSnap } from './post';
import { updateSnap } from './put';
import { deleteSnap } from './delete';

export const SNAPS_URL = `${API_URL}/snapdex`;

export interface Snap {
	id: string;
	user_id: string;
	media_url: string;
	caption: string;
	tags: string[];
	created_at: string;
	updated_at: string;
	deleted_at?: string;
}

export { getAllSnaps, getMySnaps, getSnap, createSnap, updateSnap, deleteSnap };
