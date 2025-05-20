<script lang="ts">
	import { fetchProfile } from "$lib/api/auth";
	import { user } from "$lib/store/user";
	import { onMount } from "svelte";

	// You can import API functions if needed
	// import { fetchUserStats, fetchBlogStats } from '$lib/api';
	
    let username: string | undefined = $user?.username

    onMount(async () => {
        const userData = await fetchProfile()
        user.set(userData.user)
        username = userData.user.username
    })
</script>

<svelte:head>
	<title>Dashboard</title>
</svelte:head>

<section class="p-6 space-y-6">
	<h1 class="text-3xl font-bold">Welcome back, {username || 'user'}!</h1>

	<div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
		<div class="bg-white rounded-2xl shadow p-4">
			<p class="text-gray-600">Total Users</p>
			<p class="text-2xl font-semibold">123</p>
		</div>
		<div class="bg-white rounded-2xl shadow p-4">
			<p class="text-gray-600">Blog Posts</p>
			<p class="text-2xl font-semibold">42</p>
		</div>
		<div class="bg-white rounded-2xl shadow p-4">
			<p class="text-gray-600">New Signups</p>
			<p class="text-2xl font-semibold">8 today</p>
		</div>
	</div>

	<div class="bg-white rounded-2xl shadow p-6">
		<h2 class="text-xl font-semibold mb-4">Recent Activity</h2>
		<ul class="space-y-2 text-gray-700">
			<li>✅ User <strong>Alice</strong> signed up</li>
			<li>📝 New blog post published: <strong>“How to train your Pikachu”</strong></li>
			<li>👤 User <strong>Bob</strong> updated profile</li>
		</ul>
	</div>
</section>
