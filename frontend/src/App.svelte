<script lang="ts">
    import Section from "./lib/Section.svelte";
    import configStore from './stores/config.store';
    import {fetchAsset} from "./services/assets.service";

    const [config, loading, error] = configStore();
</script>

<svelte:head>
    {#await fetchAsset('themes/default.css') then res}
        {#if (res.ok)}
            <link href="{res.url}" rel="stylesheet"/>
        {/if}
    {/await}

    {#if $config}
        <title>{$config.title}</title>
    {/if}
</svelte:head>

<main class="app-main">
    {#if $loading}
        Loading...
    {:else if $error}
        Error: {$error}
    {:else}
        <h1 class="app-title">{$config.title || 'HEDYWYD?Y!'}</h1>

        <div class="app-sections-container-wrapper">
            {#each $config.sections || [] as section}
                <Section section={section}/>
            {/each}
        </div>
    {/if}
</main>

<style>
</style>
