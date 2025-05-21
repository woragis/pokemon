import { writable } from 'svelte/store';
import type { User } from '$lib/types/user';
import { getCookie, setCookie } from '$lib/api';

export const user = writable<User | null>(null);

export function getUser() {
	const userCookie: string | null = getCookie('user');
	user.set(userCookie ? JSON.parse(userCookie) : null);
}

export function setUser(data: User) {
	user.set(data);
	setCookie('user', JSON.stringify(data));
}
