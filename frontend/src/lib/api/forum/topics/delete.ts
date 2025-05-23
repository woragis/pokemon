import { getHeaders } from '../..';
import { FORUM_TOPICS_URL } from '.';

export interface DeleteForumTopicProps {
	id: string;
}

export async function deleteForumTopic({ id }: DeleteForumTopicProps) {
	const headers = getHeaders();
	try {
		const res = await fetch(`${FORUM_TOPICS_URL}/${id}`, {
			method: 'DELETE',
			headers
		});
		if (!res.ok) throw new Error();
		return await res.json();
	} catch {
		throw new Error('Error deleting forum topic');
	}
}
