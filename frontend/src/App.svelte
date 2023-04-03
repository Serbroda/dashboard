<script lang="ts">
    import {apiService} from "./services/api.service.js";
    import Section from "./lib/Section.svelte";

</script>

<svelte:head>
    {#await apiService.get() then data}
        <title>{data.title}</title>
    {/await}
</svelte:head>

<main class="h-full">
    <h1>Dashboard</h1>

    <div class="container">
        {#await apiService.get() then data}

            {#each data.sections || [] as section}
                <Section section={section}/>
            {/each}
        {/await}
    </div>
</main>

<style>
</style>
