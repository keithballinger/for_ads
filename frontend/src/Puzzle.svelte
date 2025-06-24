<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { Button, TextInput } from '@mantine/core'

  const dispatch = createEventDispatcher()

  let answer = ''

  function checkAnswer() {
    if (answer.toLowerCase() === 'i love you') {
      dispatch('solved')
    } else {
      alert('Wrong answer, try again!')
    }
  }
</script>

<div>
  <h1>A little puzzle for you...</h1>
  <p>What are the three words I want to say to you most?</p>
  <TextInput bind:value={answer} placeholder="Your answer" />
  <Button on:click={checkAnswer}>Submit</Button>
</div>
