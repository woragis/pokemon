export const BASE_URL = 'http://localhost:3000';
export const API_URL = `${BASE_URL}/api`;

export function getHeaders() {
	const token = '';
	return {
		Authorization: `Bearer ${token}`,
		'Content-Type': 'application/json'
	};
}
