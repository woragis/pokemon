export const BASE_URL = 'http://localhost:3000';
export const API_URL = `${BASE_URL}/api`;

export function getHeaders() {
	const token = getCookie('token');
	return {
		Authorization: `Bearer ${token}`,
		'Content-Type': 'application/json'
	};
}

export function setCookie(name: string, value: string, days = 7) {
	const expires = new Date(Date.now() + days * 864e5).toUTCString();
	document.cookie = `${encodeURIComponent(name)}=${encodeURIComponent(value)}; expires=${expires}; path=/`;
}

export function getCookie(name: string): string | null {
	const cookies = document.cookie.split('; ');
	for (const cookie of cookies) {
		const [key, value] = cookie.split('=');
		if (decodeURIComponent(key) === name) {
			return decodeURIComponent(value);
		}
	}
	return null;
}
