import { PERMISSIONS_URL, type Permission } from '.';
import { getCookie } from '..';

export async function createPermission(
	data: Omit<Permission, 'id' | 'created_at' | 'updated_at'>
): Promise<Permission> {
	const token = getCookie('token');
	const res = await fetch(`${PERMISSIONS_URL}`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`
		},
		body: JSON.stringify(data)
	});
	if (!res.ok) throw new Error('Error creating permission');
	return await res.json();
}
