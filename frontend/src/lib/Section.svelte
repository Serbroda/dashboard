<script lang="ts">
  import type {Section} from "../models/config.model";

  export let section: Section;
</script>

<div class="section--container">
  <div class="section--header">
    {section.name}
  </div>
  <div class="section--items-container">
    {#each section.items as item}
      <a class="section--item" href={item.url}>{item.name}</a>
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