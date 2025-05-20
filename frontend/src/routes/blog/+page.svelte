<script lang='ts'>
  import { onMount } from 'svelte';
  import { getPosts } from '../../lib/api';
	import type { BlogPost } from '$lib/types';

  let posts: BlogPost[] = [];

  onMount(async () => {
    posts = await getPosts();
  });
</script>

<h1>Pokémon Blog</h1>

{#if posts.length === 0}
  <p>Loading posts...</p>
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
