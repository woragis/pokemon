<script lang="ts">
	import type { Permission } from '$lib/api/permissions';
	import { getPermissions } from '$lib/api/permissions/get';
	import type { Role } from '$lib/api/roles';
	import { getRoles } from '$lib/api/roles/get';
	import { onMount } from 'svelte';
	let roles: Role[];
	let permissions: Permission[];
	let newRole = { name: '', description: '' };
	let newPermission = { name: '', description: '' };
	let error = '';

	async function handleFetch() {
		try {
			roles = await getRoles();
		} catch (err) {
			error = 'Error fetching roles';
		}
		try {
			permissions = await getPermissions();
		} catch (err) {
			error = 'Error fetching permissions';
		}
	}
	onMount(handleFetch);
</script>

{#if error}
	<p class="error">{error}</p>
{/if}

<div class="bg-gradient-to-r from-yellow-500 to-yellow-600 px-4 py-16">
	<div class="container mx-auto">
		<h1 class="mb-4 text-center text-3xl font-bold text-white md:text-4xl">
			Manage Roles & Permissions
		</h1>
		<p class="mx-auto max-w-2xl text-center text-yellow-100">
			Define and organize your platform access levels and capabilities with ease.
		</p>
	</div>
</div>

<div class="container mx-auto px-4 py-12">
	<div class="grid grid-cols-1 gap-8 lg:grid-cols-2">
		<!-- Roles Section -->
		<div class="rounded-lg bg-white p-6 shadow-md">
			<h2 class="mb-4 text-xl font-semibold text-red-600">Roles</h2>
			<form class="mb-6 space-y-4">
				<input
					type="text"
					placeholder="Role name"
					bind:value={newRole.name}
					class="w-full rounded border px-4 py-2"
				/>
				<textarea
					placeholder="Description"
					bind:value={newRole.description}
					class="w-full rounded border px-4 py-2"
				></textarea>
				<button class="rounded bg-red-600 px-4 py-2 font-semibold text-white hover:bg-red-700"
					>Add Role</button
				>
			</form>
			<ul class="space-y-4">
				{#each roles as role (role.name)}
					<li class="flex items-center justify-between rounded border px-4 py-2">
						<div>
							<h3 class="font-semibold text-gray-800">{role.name}</h3>
							<p class="text-sm text-gray-500">{role.description}</p>
						</div>
						<button class="text-sm text-red-500 hover:underline">Delete</button>
					</li>
				{/each}
			</ul>
		</div>

		<!-- Permissions Section -->
		<div class="rounded-lg bg-white p-6 shadow-md">
			<h2 class="mb-4 text-xl font-semibold text-blue-600">Permissions</h2>
			<form class="mb-6 space-y-4">
				<input
					type="text"
					placeholder="Permission name"
					bind:value={newPermission.name}
					class="w-full rounded border px-4 py-2"
				/>
				<textarea
					placeholder="Description"
					bind:value={newPermission.description}
					class="w-full rounded border px-4 py-2"
				></textarea>
				<button class="rounded bg-blue-600 px-4 py-2 font-semibold text-white hover:bg-blue-700"
					>Add Permission</button
				>
			</form>
			<ul class="space-y-4">
				{#each permissions as permission (permission.name)}
					<li class="flex items-center justify-between rounded border px-4 py-2">
						<div>
							<h3 class="font-semibold text-gray-800">{permission.name}</h3>
							<p class="text-sm text-gray-500">{permission.description}</p>
						</div>
						<button class="text-sm text-blue-500 hover:underline">Delete</button>
					</li>
				{/each}
			</ul>
		</div>
	</div>
</div>
