<script lang="ts">
  import Puzzle from './Puzzle.svelte'
  import Sonnet from './Sonnet.svelte'

  let puzzleSolved = false
</script>

<main>
  {#if !puzzleSolved}
    <Puzzle on:solved={() => puzzleSolved = true} />
  {:else}
    <Sonnet />
  {/if}
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
  }
</style>
