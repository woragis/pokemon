<script lang="ts">
  import { onMount } from 'svelte'
  import { Menu, X, Search, User, Home, BookOpen, Users, Newspaper, PenTool, HelpCircle, LogOut } from 'lucide-svelte'
  import Link from '../common/Link.svelte'

  let isMenuOpen = false
  let isScrolled = false
  let user = null

  onMount(() => {
    const handleScroll = () => {
      isScrolled = window.scrollY > 10
    }

    window.addEventListener('scroll', handleScroll)
    return () => window.removeEventListener('scroll', handleScroll)
  })

  // Example: Replace with your own auth integration (e.g., Supabase or Firebase)
  onMount(() => {
    // const { data: { subscription } } = supabase.auth.onAuthStateChange((event, session) => {
    //   user = session?.user || null
    // })
    // return () => subscription.unsubscribe()
  })

  const handleSignOut = async () => {
    // await supabase.auth.signOut()
    // window.location.href = '/'
  }
</script>

<header class={`fixed w-full z-50 transition-all duration-300 ${isScrolled ? 'bg-white bg-opacity-95 shadow-md' : 'bg-transparent'}`}>
  <div class="container mx-auto px-4">
    <div class="flex items-center justify-between h-16">
      <div class="flex items-center">
        <div class="text-2xl font-bold text-red-600 flex items-center gap-2">
          <div class="w-8 h-8 rounded-full bg-red-600 flex items-center justify-center">
            <div class="w-3 h-3 bg-white rounded-full"></div>
          </div>
          <span>PokéHub</span>
        </div>
      </div>

      <div class="hidden md:flex items-center space-x-8">
        <Link to="/" icon={Home} label="Home" />
        <Link to="/pokedex" icon={BookOpen} label="Pokédex" />
        <Link to="/teams" icon={Users} label="Teams" />
        <Link to="/news" icon={Newspaper} label="News" />
        <Link to="/forum" icon={PenTool} label="Forum" />
        <Link to="/blog" icon={BookOpen} label="Blog" />
        <Link to="/faq" icon={HelpCircle} label="FAQ" />
      </div>

      <div class="hidden md:flex items-center space-x-4">
        <button class="p-2 rounded-full hover:bg-gray-100 transition-colors">
          <Search class="w-5 h-5 text-gray-600" />
        </button>

        {#if user}
          <div class="relative group">
            <a href="/profile" class="p-2 rounded-full hover:bg-gray-100 transition-colors">
              <User class="w-5 h-5 text-gray-600" />
            </a>
            <div class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 hidden group-hover:block">
              <a href="/profile" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Profile</a>
              <button on:click={handleSignOut} class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Sign Out</button>
            </div>
          </div>
        {:else}
          <a href="/auth" class="p-2 rounded-full hover:bg-gray-100 transition-colors">
            <User class="w-5 h-5 text-gray-600" />
          </a>
        {/if}
      </div>

      <div class="md:hidden flex items-center">
        <button
          class="p-2 rounded-md text-gray-600 hover:text-gray-800 hover:bg-gray-100 focus:outline-none"
          on:click={() => isMenuOpen = !isMenuOpen}
        >
          {#if isMenuOpen}
            <X class="w-6 h-6" />
          {:else}
            <Menu class="w-6 h-6" />
          {/if}
        </button>
      </div>
    </div>
  </div>

  {#if isMenuOpen}
    <div class="md:hidden bg-white shadow-lg p-4">
      <div class="flex flex-col space-y-4">
        <Link to="/" icon={Home} label="Home" mobile />
        <Link to="/pokedex" icon={BookOpen} label="Pokédex" mobile />
        <Link to="/teams" icon={Users} label="Teams" mobile />
        <Link to="/news" icon={Newspaper} label="News" mobile />
        <Link to="/forum" icon={PenTool} label="Forum" mobile />
        <Link to="/blog" icon={BookOpen} label="Blog" mobile />
        <Link to="/faq" icon={HelpCircle} label="FAQ" mobile />

        <div class="pt-4 flex items-center space-x-4 border-t border-gray-200">
          <button class="flex items-center space-x-2 text-gray-600">
            <Search class="w-5 h-5" />
            <span>Search</span>
          </button>

          {#if user}
            <a href="/profile" class="flex items-center space-x-2 text-gray-600">
              <User class="w-5 h-5" />
              <span>Profile</span>
            </a>
            <button on:click={handleSignOut} class="flex items-center space-x-2 text-gray-600">
              <LogOut class="w-5 h-5" />
              <span>Sign Out</span>
            </button>
          {:else}
            <a href="/auth" class="flex items-center space-x-2 text-gray-600">
              <User class="w-5 h-5" />
              <span>Sign In</span>
            </a>
          {/if}
        </div>
      </div>
    </div>
  {/if}
</header>
