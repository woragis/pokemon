import { getCookie } from '..';
import { PERMISSIONS_URL } from '.';

export async function deletePermission(id: string): Promise<void> {
	const token = getCookie('token');
	const res = await fetch(`${PERMISSIONS_URL}/${id}`, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`
		}
	});
	if (!res.ok) throw new Error('Error deleting permission');
}
