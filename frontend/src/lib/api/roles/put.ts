import { ROLES_URL, type Role } from '.';
import { getCookie } from '..';

export async function updateRole(id: string, data: Partial<Role>) {
	const token = getCookie('token');
	const res = await fetch(`${ROLES_URL}/${id}`, {
		method: 'PUT',
		headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${token}` },
		body: JSON.stringify(data)
	});
	if (!res.ok) throw new Error('Error updating role');
	return await res.json();
}
