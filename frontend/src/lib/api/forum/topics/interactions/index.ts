import { getCookie } from '$lib/api';
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

export async function fetchForumCommentsByTopicId({ id }: { id: string }) {
	const res = await fetch(`${FORUM_TOPICS_URL}/${id}/comments`);
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
	return res.json();
}
