<script lang="ts">
	import "./layout.css";
	import favicon from "$lib/assets/favicon.svg";
	import { onMount } from "svelte";
	import {
		connect,
		disconnect,
		wsStore,
		type WSMessage,
	} from "$lib/client/websocket";

	let connected = $state(false);

	let messages = $state<WSMessage[]>([]);
	$effect(() => {
		console.log($inspect(messages));
	});

	onMount(() => {
		connect("ws://localhost:8080/ws");

		const unsubscribe = wsStore.subscribe((value) => {
			connected = value.connected;
			messages = value.messages;
		});

		return () => {
			unsubscribe();
			disconnect();
		};
	});

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<div class="flex flex-col h-screen">
	<div
		class="w-full bg-emerald-500 dark:bg-emerald-800 p-4 flex justify-between items-center shrink-0"
	>
		<div>
			<a href="/"
				><h2
					class="text-xl font-bold hover:cursor-pointer font-stretch-150%"
				>
					Enkrypted
				</h2></a
			>
		</div>
		<div class="flex gap-5">
			{#if connected}
				<p
					class="p-2 rounded-md preset-filled-success-700-300 dark:preset-filled-success-100-900"
				>
					Status: Connected
				</p>
			{:else}
				<p
					class="p-2 rounded-md preset-filled-error-700-300 dark:preset-filled-error-100-900"
				>
					Status: Disconnected
				</p>
			{/if}
			<a href="/login" class="btn preset-filled">Login</a>
		</div>
	</div>
	<main class="flex-1 overflow-y-auto bg-neutral-300 dark:bg-neutral-900">
		{@render children()}
	</main>
</div>
