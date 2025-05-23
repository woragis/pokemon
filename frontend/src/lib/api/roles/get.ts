import { ROLES_URL, type Role } from '.';

export async function getRoles(): Promise<Role[]> {
	const role = await fetch(`${ROLES_URL}/`);
	if (!role.ok) throw new Error('Error fetching roles');
	return await role.json();
}

export async function getRole(id: string): Promise<Role> {
	const role = await fetch(`${ROLES_URL}/${id}`);
	if (!role.ok) throw new Error('Error fetching role');
	return await role.json();
}
