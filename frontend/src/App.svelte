<script lang="ts">
import {apiService} from "./services/api.service.js";
</script>

<svelte:head>
{#await apiService.get() then data}
    <title>{data.title}</title>
{/await}
</svelte:head>

<main>
  <h1>Dashboard</h1>

  <div class="container">
    {#each [...Array(10).keys()] as i}
      <section class="section">{i+1}</section>
    {/each}
  </div>
</main>

<style>

</style>
