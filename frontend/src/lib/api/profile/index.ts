import type { DefaultResponse } from '$lib/types';
import type { User } from '$lib/types/user';
import { API_URL } from '..';
import { fetchProfile } from './get';
import { putProfile } from './put';

export const PROFILE_URL = `${API_URL}/profile`;

export interface ProfileResponse extends DefaultResponse {
	user: User;
}

export { fetchProfile, putProfile };
