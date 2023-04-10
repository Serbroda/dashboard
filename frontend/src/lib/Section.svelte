<script lang="ts">
    import type {Section} from "../models/config.model";
    import SectionItem from "./SectionItem.svelte";

    export let section: Section;
</script>

<div class="p-4">
    <div class="section--header">
        {section.name}
    </div>
    <div class="grid grid-cols-4 gap-4">
        {#each section.items as item}
            <SectionItem item={item}/>
        {/each}
    </div>
</div>

<style>
    /*.section--container {*/
    /*  display: flex;*/
    /*  flex-direction: column;*/
    /*  padding: 8px;*/
    /*  border-radius: 6px;*/
    /*  background-color: #2f2f2f;*/
    /*}*/
    /*.section--header {*/
    /*}*/
    /*.section--items-container {*/
    /*  display: grid;*/
    /*  grid-template-columns: 1fr 1fr 1fr 1fr;*/
    /*  gap: 4px;*/
    /*  margin-top: 4px;*/
    /*}*/
    /*.section--item {*/
    /*  padding: 8px;*/
    /*  border-radius: 6px;*/
    /*  background-color: #1f1f1f;*/
    /*}*/
</style>