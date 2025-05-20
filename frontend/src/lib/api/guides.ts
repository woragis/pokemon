let guide = {
	id: 1,
	title: 'How to Defeat Brock',
	category: 'Gym Leaders',
	content: 'Use grass- or water-type Pokémon to beat Brock easily.',
	createdAt: new Date(),
	updatedAt: new Date()
};
type Guide = typeof guide;
let guides = [
	{
		id: 1,
		title: 'How to Defeat Brock',
		category: 'Gym Leaders',
		content: 'Use grass- or water-type Pokémon to beat Brock easily.',
		createdAt: new Date(),
		updatedAt: new Date()
	}
];

export function getGuides() {
	return guides;
}

export function getGuide(id: string) {
	return guides.find((g) => g.id === Number(id));
}

export function createGuide(guide: Omit<Guide, 'id' | 'createdAt' | 'updatedAt'>) {
	const id = guides.length + 1;
	const now = new Date();
	const newGuide = { ...guide, id, createdAt: now, updatedAt: now };
	guides.push(newGuide);
	return newGuide;
}

export function updateGuide(id: string, updates: Guide) {
	const guide = getGuide(id);
	if (guide) Object.assign(guide, updates, { updatedAt: new Date() });
	return guide;
}

export function deleteGuide(id: number) {
	guides = guides.filter((g) => g.id !== Number(id));
}
