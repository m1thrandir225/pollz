<template>
  <div v-if="status === 'pending'">Loading...</div>
  <div
    v-else-if="status === 'success' && data"
    class="max-w-[650px] my-4 flex flex-col items-start gap-8 mx-auto"
  >
    <div class="flex flex-row justify-between items-center w-full">
      <h1>{{ data.poll.description }}</h1>

      <p class="">
        Active until <br >
        <span class="text-[12px] text-orange-600">{{
          formatDate(new Date(data.poll.active_until))
        }}</span>
      </p>
    </div>
    <div v-if="isPollByUser">
      <h1>Not by user</h1>
      <form>
        <input type="radio" >
      </form>
    </div>
    <div v-else>
      <div v-for="option in data.options" :key="option.id">
        {{ option.option_text }}
        {{ option.vote_count }}
      </div>
    </div>
  </div>
  <div v-else>Error...</div>
</template>

<script setup lang="ts">
const route = useRoute();

const authStore = useAuthStore();

const isPollByUser = computed(() => {
  return data.value?.poll.created_by === authStore.user?.id;
});

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
