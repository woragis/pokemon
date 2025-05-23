import { PERMISSIONS_URL, type Permission } from '.';
import { getCookie } from '..';

export async function updatePermission(id: string, data: Partial<Permission>): Promise<Permission> {
	const token = getCookie('token');
	const res = await fetch(`${PERMISSIONS_URL}/${id}`, {
		method: 'PUT',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`
		},
		body: JSON.stringify(data)
	});
	if (!res.ok) throw new Error('Error updating permission');
	return await res.json();
}
