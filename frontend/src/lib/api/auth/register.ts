import { AUTH_URL, type AuthResponse } from '.';
import { getHeaders, setCookie } from '..';

interface RegisterProps {
	username: string;
	email: string;
	password: string;
}

export async function register({ username, email, password }: RegisterProps) {
	try {
		const res = await fetch(`${AUTH_URL}/register`, {
			method: 'POST',
			headers: getHeaders(),
			body: JSON.stringify({ username, email, password })
		});
		if (!res.ok) throw new Error('Registration failed');
		const response: AuthResponse = await res.json();
		setCookie('token', response.token);
		return response;
	} catch (e: any) {
		throw new Error(e.message || 'Error registering');
	}
}
