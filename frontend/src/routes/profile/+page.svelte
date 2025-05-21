<script lang="ts">
	import { onMount } from 'svelte';
	import { User, Edit2, Save } from 'lucide-svelte';

	interface Profile {
		username: string;
		full_name: string;
		favorite_pokemon: string;
		favorite_game: string;
		age: number | '';
	}

	const profile: Profile = {
		username: 'woragis',
		full_name: 'Jezreel',
		favorite_game: 'Red',
		favorite_pokemon: 'Charizard',
		age: 21
	};
	let isEditing = false;
	let loading = true;

	onMount(() => {
		fetchProfile();
	});

	async function fetchProfile() {
		try {
			// Example fetching logic here, replace with your actual data fetch
			// const user = await supabase.auth.getUser()
			// if (!user) return
			// const { data, error } = await supabase.from('profiles').select('*').eq('id', user.id).single()
			// if (error) throw error
			// profile = data
			// For demo, set dummy profile:
			// profile = {
			// 	username: 'pikachu_master',
			// 	full_name: 'Ash Ketchum',
			// 	favorite_pokemon: 'Pikachu',
			// 	favorite_game: 'Pokémon Red',
			// 	age: 15
			// };
		} catch (error) {
			console.error('Error fetching profile:', error);
		} finally {
			loading = false;
		}
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		try {
			// Your update logic here
			// const user = await supabase.auth.getUser()
			// if (!user) return
			// const { error } = await supabase.from('profiles').update(profile).eq('id', user.id)
			// if (error) throw error
			isEditing = false;
		} catch (error) {
			console.error('Error updating profile:', error);
		}
	}
</script>

{#if loading}
	<div class="min-h-screen bg-gray-50 pt-16">
		<div class="container mx-auto px-4 py-8">
			<div class="text-center">Loading...</div>
		</div>
	</div>
{:else}
	<div class="min-h-screen bg-gray-50 pt-16">
		<div class="container mx-auto px-4 py-8">
			<div class="mx-auto max-w-2xl overflow-hidden rounded-lg bg-white shadow-md">
				<div
					class="flex items-center justify-between bg-gradient-to-r from-red-600 to-red-700 px-6 py-4"
				>
					<h1 class="flex items-center text-2xl font-bold text-white">
						<User class="mr-2" />
						Profile
					</h1>
					<button
						on:click={() => (isEditing = !isEditing)}
						class="text-white transition-colors hover:text-red-200"
						aria-label={isEditing ? 'Save profile' : 'Edit profile'}
						type="button"
					>
						{#if isEditing}
							<Save />
						{:else}
							<Edit2 />
						{/if}
					</button>
				</div>

				<form on:submit|preventDefault={handleSubmit} class="p-6">
					<div class="space-y-6">
						<div>
							<label class="block text-sm font-medium text-gray-700">Username</label>
							<input
								type="text"
								bind:value={profile.username}
								disabled={!isEditing}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500"
								required
							/>
						</div>

						<div>
							<label class="block text-sm font-medium text-gray-700">Full Name</label>
							<input
								type="text"
								bind:value={profile.full_name}
								disabled={!isEditing}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500"
								required
							/>
						</div>

						<div>
							<label class="block text-sm font-medium text-gray-700">Favorite Pokémon</label>
							<input
								type="text"
								bind:value={profile.favorite_pokemon}
								disabled={!isEditing}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500"
							/>
						</div>

						<div>
							<label class="block text-sm font-medium text-gray-700">Favorite Game</label>
							<input
								type="text"
								bind:value={profile.favorite_game}
								disabled={!isEditing}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500"
							/>
						</div>

						<div>
							<label class="block text-sm font-medium text-gray-700">Age</label>
							<input
								type="number"
								bind:value={profile.age}
								disabled={!isEditing}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-red-500 focus:ring-red-500"
								min="0"
							/>
						</div>

						{#if isEditing}
							<div class="flex justify-end">
								<button
									type="submit"
									class="rounded-md bg-red-600 px-4 py-2 text-white transition-colors hover:bg-red-700"
								>
									Save Changes
								</button>
							</div>
						{/if}
					</div>
				</form>
			</div>
		</div>
	</div>
{/if}
