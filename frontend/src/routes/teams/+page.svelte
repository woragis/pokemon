<script>
	import TeamBuilder from '$lib/components/teams/TeamBuilder.svelte';
	import { Shield, AlertTriangle } from 'lucide-svelte';
</script>

<div class="header-gradient padding-x padding-y-large">
	<div class="container">
		<h1 class="title-main">Team Builder</h1>
		<p class="subtitle">
			Create the perfect Pokémon team. Analyze type advantages, build balanced teams, and prepare
			for battles.
		</p>
	</div>
</div>

<div class="padding-x padding-y-medium container">
	<div class="grid-layout">
		<div class="main-section">
			<TeamBuilder />
		</div>

		<div class="sidebar">
			<div class="card margin-bottom">
				<h3 class="card-title">Team Analysis</h3>

				<div class="progress-block margin-bottom">
					<div class="progress-label">
						<span>Offensive Coverage</span>
						<span class="text-green-600">Strong</span>
					</div>
					<div class="progress-bar-bg">
						<div class="progress-bar-green" style="width: 75%;"></div>
					</div>
				</div>

				<div class="progress-block margin-bottom">
					<div class="progress-label">
						<span>Defensive Coverage</span>
						<span class="text-yellow-600">Average</span>
					</div>
					<div class="progress-bar-bg">
						<div class="progress-bar-yellow" style="width: 50%;"></div>
					</div>
				</div>

				<div class="progress-block margin-bottom">
					<div class="progress-label">
						<span>Type Diversity</span>
						<span class="text-green-600">Excellent</span>
					</div>
					<div class="progress-bar-bg">
						<div class="progress-bar-green" style="width: 90%;"></div>
					</div>
				</div>

				<div class="weaknesses-section margin-top-large">
					<h4 class="section-title">
						<Shield size="16" class="icon-blue" />
						Weaknesses
					</h4>
					<div class="small-text text-gray-700">
						<p class="margin-bottom-small">Your team has weaknesses against:</p>
						<ul class="padded-left small-text list-disc">
							<li>Ground-type attacks</li>
							<li>Ice-type attacks</li>
						</ul>
					</div>
				</div>

				<div class="suggestions-section margin-top-medium">
					<h4 class="section-title">
						<AlertTriangle size="16" class="icon-yellow" />
						Suggestions
					</h4>
					<div class="small-text text-gray-700">
						<p>Consider adding a Pokémon with resistance to Ground and Ice attacks.</p>
					</div>
				</div>
			</div>

			<div class="card">
				<h3 class="card-title">Saved Teams</h3>
				<div class="saved-teams-list">
					<div class="saved-team">
						<h4 class="saved-team-title">Elite Four Challenge</h4>
						<p class="saved-team-date">Last updated: 2 days ago</p>
					</div>
					<div class="saved-team">
						<h4 class="saved-team-title">Water Type Team</h4>
						<p class="saved-team-date">Last updated: 1 week ago</p>
					</div>
					<div class="saved-team">
						<h4 class="saved-team-title">Championship Team</h4>
						<p class="saved-team-date">Last updated: 2 weeks ago</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	/* Containers */
	.container {
		max-width: 1024px;
		margin-left: auto;
		margin-right: auto;
	}

	/* Spacing */
	.padding-x {
		padding-left: 1rem; /* 16px */
		padding-right: 1rem;
	}
	.padding-y-large {
		padding-top: 4rem; /* 64px */
		padding-bottom: 4rem;
	}
	.padding-y-medium {
		padding-top: 3rem; /* 48px */
		padding-bottom: 3rem;
	}
	.margin-bottom {
		margin-bottom: 1.5rem; /* 24px */
	}
	.margin-bottom-small {
		margin-bottom: 0.25rem; /* 4px */
	}
	.margin-top-large {
		margin-top: 1.5rem; /* 24px */
	}
	.margin-top-medium {
		margin-top: 1rem; /* 16px */
	}
	.margin-bottom-small {
		margin-bottom: 0.25rem;
	}
	.margin-bottom {
		margin-bottom: 1.5rem;
	}
	.margin-top {
		margin-top: 1rem;
	}

	/* Typography */
	.title-main {
		margin-bottom: 1rem; /* 16px */
		text-align: center;
		font-size: 1.875rem; /* 30px */
		font-weight: 700;
		color: white;
	}
	@media (min-width: 768px) {
		.title-main {
			font-size: 2.25rem; /* 36px */
		}
	}
	.subtitle {
		margin-left: auto;
		margin-right: auto;
		margin-bottom: 1.5rem;
		max-width: 42rem; /* 672px */
		text-align: center;
		color: #bfdbfe; /* Tailwind blue-100 */
	}
	.card-title {
		margin-bottom: 1rem;
		font-size: 1.125rem; /* 18px */
		font-weight: 600;
		color: #111827; /* gray-900 */
	}
	.section-title {
		margin-bottom: 0.5rem;
		display: flex;
		align-items: center;
		font-size: 0.875rem; /* 14px */
		font-weight: 600;
		color: #111827;
	}
	.small-text {
		font-size: 0.875rem;
	}
	.text-gray-700 {
		color: #4b5563;
	}
	.text-green-600 {
		color: #16a34a;
	}
	.text-yellow-600 {
		color: #ca8a04;
	}
	.text-white {
		color: white;
	}
	.saved-team-title {
		font-weight: 500;
	}
	.saved-team-date {
		margin-top: 0.25rem;
		font-size: 0.75rem;
		color: #6b7280;
	}

	/* Colors */
	.bg-white {
		background-color: white;
	}
	.bg-gray-50 {
		background-color: #f9fafb;
	}
	.bg-gray-200 {
		background-color: #e5e7eb;
	}
	.bg-green-500 {
		background-color: #22c55e;
	}
	.bg-yellow-500 {
		background-color: #eab308;
	}

	/* Gradient background */
	.header-gradient {
		background: linear-gradient(to right, #2563eb, #1d4ed8); /* blue-600 to blue-700 */
	}

	/* Layout */
	.grid-layout {
		display: grid;
		grid-template-columns: 1fr;
		gap: 2rem; /* 32px */
	}
	@media (min-width: 1024px) {
		.grid-layout {
			grid-template-columns: repeat(4, 1fr);
		}
	}
	.main-section {
		grid-column: span 3;
	}
	.sidebar {
		grid-column: span 1;
	}

	/* Card */
	.card {
		border-radius: 0.5rem; /* 8px */
		background-color: white;
		padding: 1.5rem; /* 24px */
		box-shadow: 0 1px 2px rgb(0 0 0 / 0.1);
	}

	/* Progress bar */
	.progress-block {
		margin-bottom: 1rem;
	}
	.progress-label {
		margin-bottom: 0.5rem;
		display: flex;
		justify-content: space-between;
		font-size: 0.875rem;
		font-weight: 500;
	}
	.progress-bar-bg {
		height: 0.5rem; /* 8px */
		width: 100%;
		background-color: #e5e7eb;
		border-radius: 9999px;
	}
	.progress-bar-green {
		height: 0.5rem;
		border-radius: 9999px;
		background-color: #22c55e;
	}
	.progress-bar-yellow {
		height: 0.5rem;
		border-radius: 9999px;
		background-color: #eab308;
	}

	/* Lists */
	.list-disc {
		list-style-type: disc;
	}
	.padded-left {
		padding-left: 1.25rem;
	}

	/* Icons */
	.icon-blue {
		color: #2563eb;
		margin-right: 0.25rem;
	}
	.icon-yellow {
		color: #ca8a04;
		margin-right: 0.25rem;
	}

	/* Saved Teams */
	.saved-teams-list > div:not(:last-child) {
		margin-bottom: 0.75rem;
	}
	.saved-team {
		background-color: #f9fafb;
		border-radius: 0.375rem;
		padding: 0.75rem;
	}
</style>
