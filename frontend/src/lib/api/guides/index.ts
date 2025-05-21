import { fetchGuides, fetchGuideById } from './get';
import { postGuide } from './post';
import { putGuide } from './put';
import { deleteGuide } from './delete';
import { API_URL } from '..';

export const GUIDE_URL = `${API_URL}/guides`;

export { fetchGuides, fetchGuideById, postGuide, putGuide, deleteGuide };
