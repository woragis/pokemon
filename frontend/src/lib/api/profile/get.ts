import { setUser } from '$lib/store/user';
import { PROFILE_URL, type ProfileResponse } from '.';
import { getHeaders } from '..';

export async function fetchProfile() {
	try {
		const res = await fetch(`${PROFILE_URL}/`, {
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
