import { AUTH_URL, type AuthResponse } from '.';
import { API_URL, getHeaders, setCookie } from '..';

interface RegisterProps {
	username: string;
	email: string;
	password: string;
	first_name: string;
	last_name: string;
}

export async function register({
	username,
	email,
	password,
	first_name,
	last_name
}: RegisterProps) {
	try {
		const res = await fetch(`${API_URL}/users/`, {
			method: 'POST',
			headers: getHeaders(),
			body: JSON.stringify({ username, email, password, first_name, last_name })
		});
		if (!res.ok) throw new Error('Registration failed');
		const response: AuthResponse = await res.json();
		setCookie('token', response.token);
		return response;
	} catch (e: any) {
		throw new Error(e.message || 'Error registering');
	}
}
