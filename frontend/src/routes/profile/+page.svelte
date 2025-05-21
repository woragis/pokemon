<script lang="ts">
	import { onMount } from 'svelte';
	import { User, Edit2, Save } from 'lucide-svelte';
	import { fetchProfile, putProfile } from '$lib/api/profile';
	import { user } from '$lib/store/user';
	import type { PokemonGame } from '$lib/types/games';
	import { fetchGames } from '$lib/api/games';

	let error = '';
	let age = 21;
	let isEditing = false;
	let loading = true;
	let games: PokemonGame[] = [];

	onMount(async () => {
		try {
			await fetchProfile();
			games = await fetchGames();
			loading = false;
		} catch (err) {
			error = 'Error loading profile';
			console.error(err);
		}
	});

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		try {
			if ($user) {
				await putProfile({
					user: {
						username: $user.username,
						name: $user.name,
						email: $user.email,
						favorite_pokemon: $user.favorite_pokemon,
						favorite_game_id: $user.favorite_game?.id,
						dob: ''
					}
				});
			}
			isEditing = false;
		} catch (error) {
			console.error('Error updating profile:', error);
		}
	}
</script>

{#if error}
	<p class="error">{error}</p>
{/if}
{#if loading}
	<div class="min-h-screen bg-gray-50 pt-16">
		<div class="container mx-auto px-4 py-8 text-center">Loading...</div>
	</div>
{:else if $user}
	<div class="min-h-screen bg-gray-50 pt-16">
		<div class="container mx-auto px-4 py-8">
			<div class="mx-auto max-w-2xl overflow-hidden rounded-lg bg-white shadow-md">
				<div
					class="flex items-center justify-between bg-gradient-to-r from-red-600 to-red-700 px-6 py-4"
				>
					<h1 class="flex items-center text-2xl font-bold text-white">
						<User class="mr-2" /> Profile
					</h1>
					<button
						on:click={() => (isEditing = !isEditing)}
						class="text-white transition-colors hover:text-red-200"
						type="button"
						aria-label={isEditing ? 'Save profile' : 'Edit profile'}
					>
						{#if isEditing}
							<Save />
						{:else}
							<Edit2 />
						{/if}
					</button>
				</div>

				<form on:submit|preventDefault={handleSubmit} class="space-y-6 p-6">
					<!-- Username -->
					<div>
						<label for="username" class="block text-sm font-medium text-gray-700">Username</label>
						<input
							id="username"
							type="text"
							bind:value={$user.username}
							disabled={!isEditing}
							required
							class="input"
						/>
					</div>

					<!-- Name -->
					<div>
						<label for="name" class="block text-sm font-medium text-gray-700">Name</label>
						<input
							id="name"
							type="text"
							bind:value={$user.name}
							disabled={!isEditing}
							required
							class="input"
						/>
					</div>

					<!-- Email -->
					<div>
						<label for="email" class="block text-sm font-medium text-gray-700">Email</label>
						<input
							id="email"
							type="email"
							bind:value={$user.email}
							disabled={!isEditing}
							required
							class="input"
						/>
					</div>

					<!-- Favorite Pokémon -->
					<div>
						<label for="favorite-pokemon" class="block text-sm font-medium text-gray-700">
							Favorite Pokémon
						</label>
						<input
							id="favorite-pokemon"
							type="text"
							bind:value={$user.favorite_pokemon}
							disabled={!isEditing}
							class="input"
						/>
					</div>

					<!-- Favorite Game -->
					<div>
						<label for="favorite-game" class="block text-sm font-medium text-gray-700"
							>Favorite Game</label
						>
						<select
							id="favorite-game"
							bind:value={$user.favorite_game.id}
							disabled={!isEditing}
							class="input"
						>
							<option value="">Select a game</option>
							{#each games as game}
								<option value={game.id}>{game.name}</option>
							{/each}
						</select>
					</div>

					<!-- Age -->
					<div>
						<label for="age" class="block text-sm font-medium text-gray-700">Age</label>
						<input
							id="age"
							type="number"
							bind:value={age}
							disabled={!isEditing}
							class="input"
							min="1"
							max="99"
						/>
					</div>

					<!-- Role (if not just 'user') -->
					{#if $user.role !== 'user'}
						<div>
							<label for="role" class="block text-sm font-medium text-gray-700">Role</label>
							<input id="role" type="text" value={$user.role} disabled class="input bg-gray-100" />
						</div>
					{/if}

					<!-- Metadata -->
					<div class="border-t pt-4 text-sm text-gray-500">
						<p><strong>Joined:</strong> {$user.created_at}</p>
						<p><strong>Last Updated:</strong> {$user.updated_at}</p>
						{#if $user.deleted_at}
							<p class="text-red-500"><strong>Deleted At:</strong> {$user.deleted_at}</p>
						{/if}
					</div>

					<!-- Save button -->
					{#if isEditing}
						<div class="flex justify-end pt-4">
							<button
								type="submit"
								class="rounded-md bg-red-600 px-4 py-2 text-white hover:bg-red-700"
							>
								Save Changes
							</button>
						</div>
					{/if}
				</form>
			</div>
		</div>
	</div>
{/if}

<!-- <style lang="tailwind">
		mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500;
</style> -->
