import type { User } from './user';

export interface Shout {
	id: string;
	user_id: string;
	user: User;
	content: string;
	created_at: string;
	updated_at: string;
	reshout_of_id?: string;
	reshout_of?: Shout;
	quote_content?: string;
	likes: any[];
	comments: any[];
	is_flagged: boolean;
}
