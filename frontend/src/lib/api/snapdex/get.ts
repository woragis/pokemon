import { SNAPS_URL, type Snap } from '.';
import { getCookie } from '..';

// ✅ Get All Snaps (Public)
export async function getAllSnaps(): Promise<Snap[]> {
	const res = await fetch(SNAPS_URL);
	if (!res.ok) throw new Error('Error fetching snaps');
	return await res.json();
}

// ✅ Get Single Snap by ID
export async function getSnap(id: string): Promise<Snap> {
	const res = await fetch(`${SNAPS_URL}/${id}`);
	if (!res.ok) throw new Error('Error fetching snap');
	return await res.json();
}

// ✅ Optional: Get Snaps of Authenticated User
export async function getMySnaps(): Promise<Snap[]> {
	const token = getCookie('token');
	const res = await fetch(`${SNAPS_URL}/me`, {
		headers: { Authorization: `Bearer ${token}` }
	});
	if (!res.ok) throw new Error('Error fetching your snaps');
	return await res.json();
}
