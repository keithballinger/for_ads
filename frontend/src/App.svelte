<script lang="ts">
  import { MantineProvider } from '@mantine/core';
  import Puzzle from './Puzzle.svelte'
  import Sonnet from './Sonnet.svelte'

  let puzzleSolved = false
</script>

<MantineProvider>
  <main>
    {#if !puzzleSolved}
      <Puzzle on:solved={() => puzzleSolved = true} />
    {:else}
      <Sonnet />
    {/if}
  </main>
</MantineProvider>

<style>
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
  }
</style>
