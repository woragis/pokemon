import { API_URL } from '$lib/api';

export const FORUM_TOPICS_URL = `${API_URL}/forum/topics`;

export * from './get';
export * from './post';
export * from './put';
export * from './delete';
