<script lang="ts">
  let count: number = $state(1)
  let doubleCount =  $derived(()=>count*2)
  const increment = () => {
    count += 1
  }
</script>

<button onclick={increment}>
  count is {count}
</button>
