<script>
	export let snap;

	// For fallback or defaults
	const profileImage = snap.author?.profile_image || '/default-avatar.png';
	const username = snap.author?.username || 'Unknown Trainer';
</script>

<article class="feed-card">
	<header class="feed-header">
		<img src={profileImage} alt="{username}'s avatar" class="avatar" />
		<span class="username">{username}</span>
	</header>

	<img src={snap.media_url} alt="Snap image" class="snap-image" />

	<footer class="feed-footer">
		<div class="actions">
			<button aria-label="Like">❤️</button>
			<button aria-label="Comment">💬</button>
			<button aria-label="Share">🔗</button>
		</div>
		<p class="caption">{snap.caption}</p>
	</footer>
</article>

<style>
	.feed-card {
		background: #fff;
		border-radius: 0.75rem;
		overflow: hidden;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
		display: flex;
		flex-direction: column;
		gap: 0;
	}

	.feed-header {
		display: flex;
		align-items: center;
		padding: 0.75rem 1rem;
		gap: 0.75rem;
	}

	.feed-header .avatar {
		width: 36px;
		height: 36px;
		border-radius: 9999px;
		object-fit: cover;
	}

	.feed-header .username {
		font-weight: 600;
		color: #1e293b; /* gray-800 */
	}

	.snap-image {
		width: 100%;
		height: auto;
		object-fit: cover;
		display: block;
	}

	.feed-footer {
		padding: 0.75rem 1rem;
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.feed-footer .actions {
		display: flex;
		gap: 1rem;
		font-size: 1.25rem;
	}

	.feed-footer .actions button {
		background: none;
		border: none;
		cursor: pointer;
		transition: transform 0.2s ease;
	}

	.feed-footer .actions button:hover {
		transform: scale(1.2);
	}

	.feed-footer .caption {
		font-size: 0.95rem;
		color: #374151; /* gray-700 */
	}
</style>
