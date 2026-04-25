<script lang="ts">
	import "./layout.css";
	import favicon from "$lib/assets/favicon.svg";

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<div
	class="w-full bg-neutral-200 dark:bg-neutral-900 p-4 flex justify-between items-center"
>
	<a href="/"
		><h2 class="text-xl font-bold hover:cursor-pointer">Enkrypted</h2></a
	>
	<a href="/login" class="btn preset-filled">Login</a>
</div>
{@render children()}
