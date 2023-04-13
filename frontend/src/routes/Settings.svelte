<script lang="ts">
    import configStore, {saveConfig} from "../stores/config.store";

    const [config, loading, error, get] = configStore;
    let json;

    config.subscribe(r => json = JSON.stringify($config, undefined, 2))

    function save() {
        saveConfig(json).then(() => get())
    }
</script>

<h1>Settings</h1>

<textarea class="w-full min-h-full" rows="10" bind:value={json}></textarea>

<button on:click={() => save()} class="btn">
Save
</button>
