import { getHeaders } from '../..';
import { FORUM_TOPICS_URL } from '.';
import type { ForumTopic } from '$lib/types/forum';

export interface PostForumTopicProps {
	topic: Omit<ForumTopic, 'id'>;
}

export async function postForumTopic({ topic }: PostForumTopicProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${FORUM_TOPICS_URL}/`, {
			method: 'POST',
			headers,
			body: JSON.stringify(topic)
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error creating forum topic');
	}
}
