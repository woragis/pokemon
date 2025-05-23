import { ROLES_URL, type Role } from '.';
import { getCookie } from '..';

export async function createRole(
	data: Omit<Role, 'id' | 'created_at' | 'updated_at'>
): Promise<Role> {
	const token = getCookie('token');
	const res = await fetch(`${ROLES_URL}`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${token}` },
		body: JSON.stringify(data)
	});
	if (!res.ok) throw new Error('Error creating role');
	return await res.json();
}
