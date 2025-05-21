import { fetchBlogPosts, fetchBlogPostById } from './get';
import { postBlogPost } from './post';
import { putBlogPost } from './put';
import { deleteBlogPost } from './delete';
import { API_URL } from '..';
export const BLOG_URL = `${API_URL}/blog`;
export { fetchBlogPosts, fetchBlogPostById, postBlogPost, putBlogPost, deleteBlogPost };
