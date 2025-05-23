import { FORUM_TOPICS_URL } from '.';

export interface FetchForumTopicProps {
	id: string;
}

export async function fetchForumTopics() {
	try {
		const res = await fetch(`${FORUM_TOPICS_URL}/`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error fetching forum topics');
	}
}

export async function fetchForumTopicById({ id }: FetchForumTopicProps) {
	try {
		const res = await fetch(`${FORUM_TOPICS_URL}/${id}`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error fetching forum topic');
	}
}
