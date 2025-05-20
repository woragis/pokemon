import type { BlogPost } from './types';

export const posts: BlogPost[] = [
	{
		id: '1',
		title: 'Welcome to the Pokémon Blog!',
		date: '2025-05-20',
		excerpt: 'Here you will find all the latest news and guides.',
		content: `<p>This is the first post in our Pokémon app blog. Stay tuned for more!</p>`
	},
	{
		id: '2',
		title: 'New Generation Released',
		date: '2025-05-18',
		excerpt: 'Learn about the new Pokémon generation and features.',
		content: `<p>The new generation brings exciting gameplay mechanics and new Pokémon species!</p>`
	}
];

export function getPosts() {
	return Promise.resolve(posts);
}

export function getPostById(id: string) {
	const post = posts.find((p) => p.id === id);
	return Promise.resolve(post);
}
