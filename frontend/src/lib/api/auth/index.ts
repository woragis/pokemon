import { API_URL } from '..';
export const AUTH_URL = `${API_URL}/auth`;

export interface AuthResponse {
	token: string;
}

import { fetchProfile } from './get';
import { login, usernameLogin } from './login';
import { register } from './register';

export { fetchProfile, login, usernameLogin, register };
