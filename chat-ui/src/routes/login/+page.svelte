<script lang="ts">
    import { enhance } from "$app/forms";
</script>

<div class="w-full flex justify-center pt-8">
    <form method="POST" use:enhance class="w-full max-w-md space-y-4 p-4">
        <fieldset class="space-y-4">
            <label class="label">
                <span class="label-text">Username</span>
                <input
                    type="text"
                    name="username"
                    class="input"
                    placeholder="Username"
                />
            </label>
        </fieldset>
        <fieldset class="flex justify-end">
            <button type="submit" class="btn preset-outlined-surface-300-700"
                >Submit</button
            >
        </fieldset>
    </form>
</div>
