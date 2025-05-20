import { API_URL, getHeaders, setCookie } from '.';
import type { DefaultResponse } from '../types';
import type { User } from '../types/user';

const AUTH_URL = `${API_URL}/auth`;

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

export async function fetchProfile() {
	try {
		const res = await fetch(`${API_URL}/profile`, {
			headers: getHeaders()
		});
		if (!res.ok) throw new Error('Registration failed');
		const response: ProfileResponse = await res.json();
		console.log('profile res', response);
		return response;
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

interface AuthResponse {
	token: string;
}

interface ProfileResponse extends DefaultResponse {
	user: User;
}
