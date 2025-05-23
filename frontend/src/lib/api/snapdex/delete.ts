import { SNAPS_URL, type Snap } from '.';
import { getCookie } from '..';

// ✅ Delete Snap
export async function deleteSnap(id: string): Promise<void> {
	const token = getCookie('token');
	const res = await fetch(`${SNAPS_URL}/${id}`, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`
		}
	});
	if (!res.ok) throw new Error('Error deleting snap');
}
