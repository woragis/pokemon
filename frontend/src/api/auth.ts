import { API_URL, getHeaders } from '.';

const AUTH_URL = `${API_URL}/auth`;

export async function login({ email, password }: LoginProps) {
	try {
		const res = await fetch(`${AUTH_URL}/login`, {
			method: 'POST',
			headers: getHeaders(),
			body: JSON.stringify({ email, password })
		});
		if (!res.ok) throw new Error('Invalid credentials');
		return await res.json(); // expect token or user data
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
		return await res.json(); // expect token or user data
	} catch (e: any) {
		throw new Error(e.message || 'Error logging in');
	}
}

export async function register({ username, email, password }: RegisterProps) {
	try {
		const res = await fetch(`${AUTH_URL}/register`, {
			method: 'POST',
			headers: getHeaders(),
			body: JSON.stringify({ username, email, password })
		});
		if (!res.ok) throw new Error('Registration failed');
		return await res.json();
	} catch (e: any) {
		throw new Error(e.message || 'Error registering');
	}
}

interface LoginProps {
	email: string;
	password: string;
}

interface UsernameLoginProps {
	username: string;
	password: string;
}

interface RegisterProps {
	username: string;
	email: string;
	password: string;
}
