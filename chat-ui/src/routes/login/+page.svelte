<script lang="ts">
    import { send, wsStore, type WSMessage } from "$lib/client/websocket";
    import { onMount } from "svelte";

    let username = $state("");
    let connected = $state(false);

    onMount(() => {
        const unsubscribe = wsStore.subscribe((value) => {
            connected = value.connected;
        });

        return () => {
            unsubscribe();
        };
    });

    const login = () => {
        if (!connected) {
            alert("Not connected to server");
            return;
        }
        send({ kind: "enkr:room:identify", data: { username: username } });
    };
</script>

<div class="w-full flex justify-center pt-8">
    <form
        class="w-full max-w-md space-y-4 p-4"
        onsubmit={(e) => {
            e.preventDefault();
            login();
        }}
    >
        <fieldset class="space-y-4">
            <label class="label">
                <span class="label-text">Username</span>
                <input
                    bind:value={username}
                    type="text"
                    name="username"
                    class="input"
                    placeholder="Username"
                />
            </label>
        </fieldset>
        <fieldset class="flex justify-end">
            <button type="submit" class="btn preset-outlined-surface-300-700"
                >{connected ? "Login" : "Connecting..."}</button
            >
        </fieldset>
    </form>
</div>
