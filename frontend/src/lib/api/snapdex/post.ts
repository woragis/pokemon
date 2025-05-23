import { SNAPS_URL, type Snap } from '.';
import { getCookie } from '..';

// ✅ Create Snap
export async function createSnap(
	data: Omit<Snap, 'id' | 'user_id' | 'created_at' | 'updated_at'>
): Promise<Snap> {
	const token = getCookie('token');
	const res = await fetch(SNAPS_URL, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`
		},
		body: JSON.stringify(data)
	});
	if (!res.ok) throw new Error('Error creating snap');
	return await res.json();
}
