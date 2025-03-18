<template>
  <div v-if="status === 'pending'">Loading...</div>
  <div
    v-else-if="status === 'success' && data"
    class="max-w-[650px] my-4 flex flex-col items-start gap-8 mx-auto"
  >
    <div class="flex flex-row justify-between items-center w-full">
      <h1>{{ data.poll.description }}</h1>

      <p>{{ formatDate(new Date(data.poll.active_until)) }}</p>
    </div>
  </div>
  <div v-else>Error...</div>
</template>

<script setup lang="ts">
const route = useRoute();

const { data, error, status } = await useFetch(
  `/api/polls/${route.params.id}`,
  {
    method: "GET",
  },
);
if (error.value) {
  throw createError({
    statusCode: error.value.statusCode,
    fatal: true,
  });
}
</script>
