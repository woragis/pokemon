import { API_URL } from '..';
import type { Permission } from '../permissions';

export const ROLES_URL = `${API_URL}/roles`;

export interface Role {
	id: string;
	name: string;
	description: string;
	permissions: Permission[];
	created_at: string;
	updated_at: string;
	deleted_at?: string;
}
