<script lang="ts">
    import Section from "./lib/Section.svelte";
    import configStore from './stores/config.store';
    import {fetchAsset} from "./services/assets.service";

    const [config, loading, error] = configStore();
</script>

<svelte:head>
    {#await fetchAsset('custom.css') then res}
        {#if (res.ok)}
            <link href="{res.url}" rel="stylesheet"/>
        {/if}
    {/await}

    {#if $config}
        <title>{$config.title}</title>
    {/if}
</svelte:head>

<main class="main">
    {#if $loading}
        Loading...
    {:else if $error}
        Error: {$error}
    {:else}
        <h1>{$config.title || 'hedywyd?y!'}</h1>

        <div class="container">
            {#each $config.sections || [] as section}
                <Section section={section}/>
            {/each}
        </div>
    {/if}
</main>

<style>
</style>
