import { API_URL } from '..';
export const AUTH_URL = `${API_URL}/auth`;

export interface AuthResponse {
	token: string;
}

import { login, usernameLogin } from './login';
import { register } from './register';

export { login, usernameLogin, register };
