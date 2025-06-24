<script lang="ts">
  import { onMount } from 'svelte'
  import { Paper, Title, Text } from '@mantine/core'

  let sonnet = {
    title: '',
    lines: [],
  }

  onMount(async () => {
    const response = await fetch('/api/sonnet')
    sonnet = await response.json()
  })
</script>

<Paper withBorder shadow="md" p={30} mt={30} radius="md">
  <Title order={2} ta="center" mt="md" mb={50}>{sonnet.title}</Title>
  {#each sonnet.lines as line}
    <Text ta="center">{line}</Text>
  {/each}
</Paper>
