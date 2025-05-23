import { ROLES_URL } from '.';
import { getCookie } from '..';

export async function deleteRole(id: string): Promise<void> {
	const token = getCookie('token');
	const res = await fetch(`${ROLES_URL}/${id}`, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`
		}
	});
	if (!res.ok) throw new Error('Error deleting role');
}
