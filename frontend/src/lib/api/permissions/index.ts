import { API_URL } from '..';

export interface Permission {
	id: string;
	name: string;
	description: string;
	created_at: string;
	updated_at: string;
	deleted_at?: string;
}
export const PERMISSIONS_URL = `${API_URL}/permissions`;
