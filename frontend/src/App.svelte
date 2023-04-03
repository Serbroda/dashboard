<script lang="ts">
    import {apiService, baseUrl} from "./services/api.service.js";
    import Section from "./lib/Section.svelte";

    /*apiService.getCustomCss().then(res => {
        console.log(res)
        if (res.ok) {
            const heads = document.getElementsByTagName("head");
            if(heads?.length > 0) {
                heads[0].append(`<link href="${res.url}" rel="stylesheet">`)
            }
        }
    });*/
</script>

<svelte:head>
    {#await apiService.getCustomCss() then res}
        {#if (res.ok)}
            <link href="{res.url}" rel="stylesheet"/>
        {/if}
    {/await}

    {#await apiService.get() then data}
        <title>{data.title}</title>
    {/await}
</svelte:head>

<main class="h-full">
    <h1>Dashboard</h1>

    <button class="btn primary">Hi there</button>

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
