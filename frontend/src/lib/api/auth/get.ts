import { setUser } from '$lib/store/user';
import type { DefaultResponse } from '$lib/types';
import type { User } from '$lib/types/user';
import { API_URL, getHeaders } from '..';

interface ProfileResponse extends DefaultResponse {
	user: User;
}

export async function fetchProfile() {
	try {
		const res = await fetch(`${API_URL}/profile`, {
			headers: getHeaders()
		});
		if (!res.ok) throw new Error('Registration failed');
		const response: ProfileResponse = await res.json();
		setUser(response.user);
		return response;
	} catch (e: any) {
		throw new Error(e.message || 'Error registering');
	}
}
