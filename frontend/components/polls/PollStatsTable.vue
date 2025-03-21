<template>
  <table class="table-none w-full md:table-fixed">
    <thead>
      <tr class="border-b">
        <th>Options</th>
        <th>Vote Count</th>
      </tr>
    </thead>
    <tbody class="text-center">
      <tr v-for="option in data" :key="option.id">
        <td>{{ option.option_text }}</td>
        <td>{{ option.vote_count }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script setup lang="ts">
import type { PollOptionWithVotes } from "~/types/poll-option";

defineProps<{
  data: PollOptionWithVotes[];
}>();
</script>
