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
	author: string;
	authorAvatar: string;
	date: string;
	replies: number;
	likes: number;
	views: number;
	category: string;
	pinned: boolean;
}
