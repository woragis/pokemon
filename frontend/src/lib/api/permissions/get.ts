import { PERMISSIONS_URL, type Permission } from '.';

export async function getPermissions(): Promise<Permission[]> {
	const res = await fetch(`${PERMISSIONS_URL}/`);
	if (!res.ok) throw new Error('Error fetching permissions');
	return await res.json();
}

export async function getPermission(id: string): Promise<Permission> {
	const res = await fetch(`${PERMISSIONS_URL}/${id}`);
	if (!res.ok) throw new Error('Error fetching permission');
	return await res.json();
}
