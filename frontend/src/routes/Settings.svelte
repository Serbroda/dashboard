<script lang="ts">
    import configStore, {saveConfig} from "../stores/config.store";

    const [config, loading, error, get] = configStore();
    let json;

    config.subscribe(r => json = JSON.stringify($config))

    function save() {
        console.log(json);
        saveConfig(json)
        get()
    }
</script>

<h1>Settings</h1>

<span>{$config?.title}</span>

<textarea class="w-full min-h-full" rows="10" bind:value={json}></textarea>

<button on:click={() => save()}>
Save
</button>