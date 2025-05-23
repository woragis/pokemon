export interface ForumCategory {
	id: string;
	name: string;
	description: string;
}

export interface ForumTopic {
	id: string;
	title: string;
	content: string;
	categoryId: string;
	authorId: string;
}
