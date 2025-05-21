<script lang="ts">
	import { onMount } from 'svelte';
	import { Plus, HelpCircle } from 'lucide-svelte';
	interface FAQ {
		id: string;
		question: string;
		answer: string;
		category: string;
	}
	let faqs: FAQ[] = [];
	let showApplicationForm = false;
	let application = { role: 'writer', reason: '' };
	onMount(() => {
		fetchFAQs();
	});
	async function fetchFAQs() {
		// Fetch FAQ data from your backend
		// Example:
		const res = await fetch('/api/faqs');
		faqs = await res.json();
	}
	async function handleApplicationSubmit(e: Event) {
		e.preventDefault();
		// Example submission logic
		const res = await fetch('/api/apply', {
			method: 'POST',
			body: JSON.stringify(application),
			headers: { 'Content-Type': 'application/json' }
		});
		if (res.ok) {
			alert('Application submitted successfully!');
			showApplicationForm = false;
			application = { role: 'writer', reason: '' };
		}
	}
</script>

<div class="bg-gradient-to-r from-purple-600 to-purple-700 px-4 py-16">
	<div class="container mx-auto">
		<h1 class="mb-4 text-center text-3xl font-bold text-white md:text-4xl">
			Frequently Asked Questions
		</h1>
		<p class="mx-auto max-w-2xl text-center text-purple-100">
			Find answers to common questions and learn how to contribute to our community.
		</p>
	</div>
</div>
<div class="container mx-auto px-4 py-12">
	<div class="grid grid-cols-1 gap-8 lg:grid-cols-3">
		<!-- FAQ List -->
		<div class="lg:col-span-2">
			<div class="space-y-6">
				{#each faqs as faq (faq.id)}
					<div class="rounded-lg bg-white p-6 shadow-md">
						<h3 class="mb-2 flex items-center text-lg font-semibold text-gray-900">
							<HelpCircle class="mr-2 h-5 w-5 text-purple-600" />
							{faq.question}
						</h3>
						<p class="text-gray-600">{faq.answer}</p>
						<div class="mt-2">
							<span
								class="inline-block rounded-full bg-purple-100 px-2 py-1 text-xs font-medium text-purple-600"
							>
								{faq.category}
							</span>
						</div>
					</div>
				{/each}
			</div>
		</div>

		<!-- Application Form -->
		<div class="lg:col-span-1">
			<div class="mb-6 rounded-lg bg-white p-6 shadow-md">
				<h2 class="mb-4 text-xl font-semibold">Join Our Team</h2>
				<p class="mb-4 text-gray-600">
					Want to contribute to our community? Apply to become a writer or moderator!
				</p>
				<button
					on:click={() => (showApplicationForm = true)}
					class="flex w-full items-center justify-center rounded-md bg-purple-600 px-4 py-2 text-white transition-colors hover:bg-purple-700"
				>
					<Plus class="mr-2 h-5 w-5" />
					Apply Now
				</button>
			</div>

			{#if showApplicationForm}
				<div class="rounded-lg bg-white p-6 shadow-md">
					<h3 class="mb-4 text-lg font-semibold">Application Form</h3>
					<form on:submit|preventDefault={handleApplicationSubmit}>
						<div class="mb-4">
							<label class="mb-1 block text-sm font-medium text-gray-700">Role</label>
							<select
								bind:value={application.role}
								class="w-full rounded-md border-gray-300 shadow-sm focus:border-purple-500 focus:ring-purple-500"
							>
								<option value="writer">Blog Writer</option>
								<option value="moderator">Forum Moderator</option>
								<option value="news">News Writer</option>
							</select>
						</div>

						<div class="mb-4">
							<label class="mb-1 block text-sm font-medium text-gray-700"
								>Why do you want to join?</label
							>
							<textarea
								bind:value={application.reason}
								rows={4}
								class="w-full rounded-md border-gray-300 shadow-sm focus:border-purple-500 focus:ring-purple-500"
								placeholder="Tell us about your experience and motivation..."
							></textarea>
						</div>

						<div class="flex justify-end">
							<button
								type="submit"
								class="rounded-md bg-purple-600 px-4 py-2 text-white transition-colors hover:bg-purple-700"
							>
								Submit Application
							</button>
						</div>
					</form>
				</div>
			{/if}
		</div>
	</div>
</div>
