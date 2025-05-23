import { SNAPS_URL, type Snap } from '.';
import { getCookie } from '..';

// ✅ Optional: Update Snap
export async function updateSnap(id: string, data: Partial<Snap>): Promise<Snap> {
	const token = getCookie('token');
	const res = await fetch(`${SNAPS_URL}/${id}`, {
		method: 'PUT',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`
		},
		body: JSON.stringify(data)
	});
	if (!res.ok) throw new Error('Error updating snap');
	return await res.json();
}
