<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import { Button, TextInput, Paper, Title, Text } from '@mantine/core'

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

<Paper withBorder shadow="md" p={30} mt={30} radius="md">
  <Title order={2} ta="center" mt="md" mb={50}>A little puzzle for you...</Title>
  <Text ta="center" mb="lg">What are the three words I want to say to you most?</Text>
  <TextInput bind:value={answer} placeholder="Your answer" required />
  <Button on:click={checkAnswer} fullWidth mt="xl">
    Submit
  </Button>
</Paper>
