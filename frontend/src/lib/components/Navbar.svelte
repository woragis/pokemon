<script lang="ts">
	import { onMount } from 'svelte';
	import {
		Menu,
		X,
		Search,
		User,
		Home,
		BookOpen,
		Users,
		Newspaper,
		PenTool,
		HelpCircle,
		LogOut,
		PersonStanding
	} from 'lucide-svelte';
	import Link from './common/Link.svelte';

	let isMenuOpen = false;
	let isScrolled = false;
	let user = null;

	onMount(() => {
		const handleScroll = () => {
			isScrolled = window.scrollY > 10;
		};

		window.addEventListener('scroll', handleScroll);
		return () => window.removeEventListener('scroll', handleScroll);
	});

	// Example: Replace with your own auth integration (e.g., Supabase or Firebase)
	onMount(() => {
		// const { data: { subscription } } = supabase.auth.onAuthStateChange((event, session) => {
		//   user = session?.user || null
		// })
		// return () => subscription.unsubscribe()
	});

	const handleSignOut = async () => {
		// await supabase.auth.signOut()
		// window.location.href = '/'
	};
</script>

<header
	class={`fixed z-50 w-full transition-all duration-300 ${isScrolled ? 'bg-white bg-opacity-95 shadow-md' : 'bg-transparent'}`}
>
	<div class="container mx-auto px-4">
		<div class="flex h-16 items-center justify-between">
			<div class="flex items-center">
				<div class="flex items-center gap-2 text-2xl font-bold text-red-600">
					<div class="flex h-8 w-8 items-center justify-center rounded-full bg-red-600">
						<div class="h-3 w-3 rounded-full bg-white"></div>
					</div>
					<span>PokéHub</span>
				</div>
			</div>

			<div class="hidden items-center space-x-8 md:flex">
				<Link to="/" icon={Home} label="Home" />
				<Link to="/pokedex" icon={BookOpen} label="Pokédex" />
				<Link to="/shouts" icon={PersonStanding} label="Shouts" />
				<Link to="/teams" icon={Users} label="Teams" />
				<Link to="/news" icon={Newspaper} label="News" />
				<Link to="/forum" icon={PenTool} label="Forum" />
				<Link to="/blog" icon={BookOpen} label="Blog" />
				<Link to="/faq" icon={HelpCircle} label="FAQ" />
			</div>

			<div class="hidden items-center space-x-4 md:flex">
				<button class="rounded-full p-2 transition-colors hover:bg-gray-100">
					<Search class="h-5 w-5 text-gray-600" />
				</button>

				{#if user}
					<div class="group relative">
						<a href="/profile" class="rounded-full p-2 transition-colors hover:bg-gray-100">
							<User class="h-5 w-5 text-gray-600" />
						</a>
						<div
							class="absolute right-0 mt-2 hidden w-48 rounded-md bg-white py-1 shadow-lg group-hover:block"
						>
							<a href="/profile" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
								>Profile</a
							>
							<button
								on:click={handleSignOut}
								class="block w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-100"
								>Sign Out</button
							>
						</div>
					</div>
				{:else}
					<a href="/auth" class="rounded-full p-2 transition-colors hover:bg-gray-100">
						<User class="h-5 w-5 text-gray-600" />
					</a>
				{/if}
			</div>

			<div class="flex items-center md:hidden">
				<button
					class="rounded-md p-2 text-gray-600 hover:bg-gray-100 hover:text-gray-800 focus:outline-none"
					on:click={() => (isMenuOpen = !isMenuOpen)}
				>
					{#if isMenuOpen}
						<X class="h-6 w-6" />
					{:else}
						<Menu class="h-6 w-6" />
					{/if}
				</button>
			</div>
		</div>
	</div>

	{#if isMenuOpen}
		<div class="bg-white p-4 shadow-lg md:hidden">
			<div class="flex flex-col space-y-4">
				<Link to="/" icon={Home} label="Home" mobile />
				<Link to="/pokedex" icon={BookOpen} label="Pokédex" mobile />
				<Link to="/shouts" icon={PersonStanding} label="Shouts" mobile />
				<Link to="/teams" icon={Users} label="Teams" mobile />
				<Link to="/news" icon={Newspaper} label="News" mobile />
				<Link to="/forum" icon={PenTool} label="Forum" mobile />
				<Link to="/blog" icon={BookOpen} label="Blog" mobile />
				<Link to="/faq" icon={HelpCircle} label="FAQ" mobile />

				<div class="flex items-center space-x-4 border-t border-gray-200 pt-4">
					<button class="flex items-center space-x-2 text-gray-600">
						<Search class="h-5 w-5" />
						<span>Search</span>
					</button>

					{#if user}
						<a href="/profile" class="flex items-center space-x-2 text-gray-600">
							<User class="h-5 w-5" />
							<span>Profile</span>
						</a>
						<button on:click={handleSignOut} class="flex items-center space-x-2 text-gray-600">
							<LogOut class="h-5 w-5" />
							<span>Sign Out</span>
						</button>
					{:else}
						<a href="/auth" class="flex items-center space-x-2 text-gray-600">
							<User class="h-5 w-5" />
							<span>Sign In</span>
						</a>
					{/if}
				</div>
			</div>
		</div>
	{/if}
</header>
