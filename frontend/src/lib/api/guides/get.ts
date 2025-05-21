import { GUIDE_URL } from '.';

export interface FetchGuideProps {
	id: string;
}

export async function fetchGuides() {
	try {
		const res = await fetch(`${GUIDE_URL}/`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching guides');
	}
}

export async function fetchGuideById({ id }: FetchGuideProps) {
	try {
		const res = await fetch(`${GUIDE_URL}/${id}`);
		if (!res.ok) throw new Error();
		return await res.json();
	} catch (e: any) {
		throw new Error('Error fetching guide');
	}
}
