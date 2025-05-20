<script lang='ts'>
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { getPostById } from '../../../lib/api';
	import type { BlogPost } from '$lib/types';

  let post: BlogPost | undefined;
  let id;

  $: id = $page.params.id;

  onMount(async () => {
    post = await getPostById(id);
  });
</script>

{#if !post}
  <p>Loading post...</p>
{:else}
  <article>
    <h1>{post.title}</h1>
    <small>{post.date}</small>
    <div>{@html post.content}</div>
    <a href="/blog">← Back to Blog</a>
  </article>
{/if}
