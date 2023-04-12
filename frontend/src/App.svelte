<script lang="ts">
    import {fetchAsset} from "./services/assets.service";
    import Router from 'svelte-spa-router'
    import Home from "./routes/Home.svelte";
    import NotFound from "./routes/NotFound.svelte";
    import Settings from "./routes/Settings.svelte";
    import configStore from "./stores/config.store";
    import NavBar from "./lib/NavBar.svelte";
    import Section from "./lib/Section.svelte";

    const [config, loading, error] = configStore;

    const routes = {
        '/': Home,
        '/settings': Settings,
        '*': NotFound,
    }
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

{#if $loading}
    Loading...
{:else if $error}
    Error: {$error}
{:else}
    <NavBar title={$config?.title}/>
    <main class="app-main">
        <Router {routes}/>
    </main>
{/if}
