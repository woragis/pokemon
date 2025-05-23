import { getHeaders } from '../..';
import { FORUM_TOPICS_URL } from '.';
import type { ForumTopic } from '$lib/types/forum';

export interface PutForumTopicProps {
	id: string;
	topic: Omit<ForumTopic, 'id' | 'author_id'>;
}

export async function putForumTopic({ id, topic }: PutForumTopicProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${FORUM_TOPICS_URL}/${id}`, {
			method: 'PUT',
			headers,
			body: JSON.stringify(topic)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error updating forum topic');
	}
}
