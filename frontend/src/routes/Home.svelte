<script lang="ts">
    import Section from "../lib/Section.svelte";
    import configStore from "../stores/config.store";

    const [config] = configStore;
</script>

<div class="app-sections-container-wrapper">
    {#each $config?.sections || [] as section}
        <Section section={section}/>
    {/each}
</div>
