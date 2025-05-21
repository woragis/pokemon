import { AUTH_URL, type AuthResponse } from '.';
import { getHeaders, setCookie } from '..';

interface LoginProps {
	email: string;
	password: string;
}

interface UsernameLoginProps {
	username: string;
	password: string;
}

export async function login({ email, password }: LoginProps) {
	try {
		const res = await fetch(`${AUTH_URL}/login`, {
			method: 'POST',
			headers: getHeaders(),
			body: JSON.stringify({ email, password })
		});
		if (!res.ok) throw new Error('Invalid credentials');
		const response: AuthResponse = await res.json();
		setCookie('token', response.token);
		return response;
	} catch (e: any) {
		throw new Error(e.message || 'Error logging in');
	}
}

export async function usernameLogin({ username, password }: UsernameLoginProps) {
	try {
		const res = await fetch(`${AUTH_URL}/login`, {
			method: 'POST',
			headers: getHeaders(),
			body: JSON.stringify({ username, password })
		});
		if (!res.ok) throw new Error('Invalid credentials');
		const response: AuthResponse = await res.json();
		setCookie('token', response.token);
		return response;
	} catch (e: any) {
		throw new Error(e.message || 'Error logging in');
	}
}
