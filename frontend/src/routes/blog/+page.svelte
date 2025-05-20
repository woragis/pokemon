<script lang='ts'>
  import { onMount } from 'svelte';
  import type { BlogPost } from '$lib/types';
  import { fetchBlogPosts } from '../../api/blog';

  let posts: BlogPost[] = [];
  let error: string | null = null;
  let loading = false;

  async function handleFetch() {
    loading = true;
    try {
      posts = await fetchBlogPosts()
    } catch (e: any) {
      error = e.message;
    } finally {
      loading = false;
    }
  }

  onMount(() => {
    handleFetch();
  });
</script>

<h1>Pokémon Blog</h1>

{#if loading}
  <p>Loading blog posts...</p>
{:else if error}
  <p>Error: {error}</p>
{:else}
  <ul>
    {#each posts as post}
      <li>
        <a href={`/blog/${post.id}`}>
          <h2>{post.title}</h2>
          <small>{post.date}</small>
          <p>{post.excerpt}</p>
        </a>
      </li>
    {/each}
  </ul>
{/if}
