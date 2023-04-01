<script lang="ts">
    import {apiService} from "./services/api.service.js";
    import Section from "./lib/Section.svelte";
</script>

<svelte:head>
    {#await apiService.get() then data}
        <title>{data.title}</title>
    {/await}
</svelte:head>

<main>
    <h1>Dashboard</h1>

    <div class="container">
        {#await apiService.get() then data}

            {#each data.sections || [] as section}
                <Section title={section.name} items={section.items}/>
            {/each}
        {/await}
    </div>
</main>

<style>
    .container {
        display: grid;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        gap: 4px;
        padding: 0px 12px;
    }
</style>
