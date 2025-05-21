<script lang="ts">
  export let to: string
  export let icon
  export let label: string
  export let mobile: boolean = false
</script>

<a href={to} class={`flex items-center ${mobile ? 'space-x-2 text-gray-600' : 'space-x-1 hover:text-red-600 transition-colors'}`}>
  <svelte:component this={icon} class="w-5 h-5" />
  <span>{label}</span>
</a>
