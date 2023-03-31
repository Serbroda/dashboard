<script lang="ts">
</script>

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
