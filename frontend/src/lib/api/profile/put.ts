import { setUser } from '$lib/store/user';
import type { User } from '$lib/types/user';
import { PROFILE_URL, type ProfileResponse } from '.';
import { getHeaders } from '..';

interface PutProfileProps {
	user: Omit<User, 'id' | 'created_at' | 'updated_at' | 'deleted_at' | 'favorite_game' | 'role'>;
}

export async function putProfile({ user }: PutProfileProps) {
	try {
		const res = await fetch(`${PROFILE_URL}/`, {
			method: 'PUT',
			headers: getHeaders(),
			body: JSON.stringify(user)
		});
		if (!res.ok) throw new Error('Profile update failed');
		const response: ProfileResponse = await res.json();
		setUser(response.user);
		return response;
	} catch (e: any) {
		throw new Error(e.message || 'Error updating profile');
	}
}
