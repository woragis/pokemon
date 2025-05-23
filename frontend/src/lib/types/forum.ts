import type { User } from './user';

export interface ForumCategory {
	id: string;
	name: string;
	color: string;
	description: string;
}

export interface ForumTopic {
	id: string;
	title: string;
	content: string;
	category_id: string;
	author_id: string;
}

export interface ForumTopicResponse {
	id: string;
	title: string;
	content: string;
	author: User;
	authorAvatar: string;
	replies: number;
	likes: number;
	views: number;
	category: ForumCategory;
	pinned: boolean;
	created_at: string;
	updated_at: string;
}

export type ForumComment = {
	id: string;
	content: string;
	user: User;
	created_at: string;
	updated_at: string;
};
