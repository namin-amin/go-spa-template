<script lang="ts">
  import { onMount } from "svelte";

  let apiresponse: string | null = null;

  let inp = "";

  onMount(() => {
    getNewData();
  });

  const getNewData = () => {
    fetch(`./api/${inp}`)
      .then(async (data) => {
        apiresponse = await data.text();
      })
      .catch((err) => {
        console.log(window.location.host);
        console.log(err);
      });
  };
</script>

<main>
  <input type="text" bind:value={inp} />
  <button on:click={getNewData}> get data </button>

  {#if apiresponse !== null}
    <h1>{apiresponse}</h1>
  {/if}
</main>

<style>
</style>
