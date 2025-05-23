import { API_URL } from '..';

export const FORUM_URL = `${API_URL}/forum`;

export {
	fetchForumCategories,
	fetchForumCategoryById,
	postForumCategory,
	putForumCategory,
	deleteForumCategory
} from './categories';

export {
	fetchForumTopics,
	fetchForumTopicById,
	postForumTopic,
	putForumTopic,
	deleteForumTopic,
	likeForumTopicById,
	viewForumTopicById,
	fetchForumCommentsByTopicId,
	createForumComment
} from './topics';
