<script lang="ts">
  import { onMount } from 'svelte'

  let sonnet = {
    title: '',
    lines: [],
  }

  onMount(async () => {
    const response = await fetch('http://localhost:8080/api/sonnet')
    sonnet = await response.json()
  })
</script>

<div>
  <h1>{sonnet.title}</h1>
  {#each sonnet.lines as line}
    <p>{line}</p>
  {/each}
</div>
