<script lang="ts">
    import {apiService} from "./services/api.service.js";
    import Section from "./lib/Section.svelte";
    import configStore from './stores/config.store';

    const [config, loading, error] = configStore();
</script>

<svelte:head>
    {#await apiService.getCustomCss() then res}
        {#if (res.ok)}
            <link href="{res.url}" rel="stylesheet"/>
        {/if}
    {/await}

    {#await apiService.getConfig() then data}
        <title>{data.title}</title>
    {/await}
</svelte:head>

<main class="main">
    <h1>hedywyd?y!</h1>

    {#if $loading}
        Loading...
    {:else if $error}
        Error: {$error}
    {:else}
        <div class="container">
            {#each $config.sections || [] as section}
                <Section section={section}/>
            {/each}
        </div>
    {/if}
</main>

<style>
</style>
