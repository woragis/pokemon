import { getCookie } from '$lib/api';
import type { ForumComment } from '$lib/types/forum';
import { FORUM_TOPICS_URL } from '..';

export async function likeForumTopicById({ id }: { id: string }) {
	const token = getCookie('token');
	const res = await fetch(`${FORUM_TOPICS_URL}/${id}/like`, {
		method: 'POST',
		headers: { Authorization: `Bearer ${token}` }
	});
	if (!res.ok) throw new Error('Failed to like topic');
	return await res.json(); // expected: { liked: true/false }
}

export async function viewForumTopicById({ id }: { id: string }) {
	const token = getCookie('token');
	const res = await fetch(`${FORUM_TOPICS_URL}/${id}/view`, {
		method: 'POST',
		headers: { Authorization: `Bearer ${token}` }
	});
	if (!res.ok) throw new Error('Failed to register view');
}

interface FetchForumCommentsByTopicIdResponse {
	comments: ForumComment[];
	pagination: {
		total: number;
		limit: number;
		offset: number;
	};
}
export async function fetchForumCommentsByTopicId({
	id,
	offset = 0,
	limit = 10
}: {
	id: string;
	offset?: number;
	limit?: number;
}): Promise<FetchForumCommentsByTopicIdResponse> {
	const res = await fetch(`${FORUM_TOPICS_URL}/${id}/comments?offset=${offset}&limit=${limit}`);
	if (!res.ok) throw new Error('Failed to fetch comments');
	return res.json();
}

export async function createForumComment({ id, content }: { id: string; content: string }) {
	const token = getCookie('token');
	const res = await fetch(`${FORUM_TOPICS_URL}/${id}/comments`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${token}` },
		body: JSON.stringify({ content })
	});
	if (!res.ok) throw new Error('Failed to create comment');
	return await res.json();
}
