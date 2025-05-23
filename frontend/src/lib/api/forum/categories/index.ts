import { API_URL } from '$lib/api';

export const FORUM_CATEGORIES_URL = `${API_URL}/forum/categories`;

export * from './get';
export * from './post';
export * from './put';
export * from './delete';
