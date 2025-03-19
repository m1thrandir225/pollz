<template>
  <div v-if="status === 'pending'">Loading...</div>
  <div
    v-else-if="status === 'success' && data"
    class="mx-auto my-4 flex max-w-[650px] flex-col items-start gap-8"
  >
    <div class="flex w-full flex-row items-center justify-between">
      <h1>{{ data.poll.description }}</h1>

      <p class="">
        Active until <br />
        <span class="text-[12px] text-orange-600">{{
          formatDate(new Date(data.poll.active_until))
        }}</span>
      </p>
    </div>
    <PollVote
      v-if="!hasVoted && !isPollByUser"
      :options="data.options"
      :refresh-page-data="refresh"
    />
    <div v-else>
      <h1>Already voted ..</h1>
      <p>stats ...</p>
    </div>
  </div>
  <div v-else>Error...</div>
</template>

<script setup lang="ts">
import PollVote from "~/components/polls/PollVote.vue";

const voteStore = useVoteStore();
const route = useRoute();

const authStore = useAuthStore();

const isPollByUser = computed(() => {
  return data.value?.poll.created_by === authStore.user?.id;
});

const hasVoted = computed(() => {
  if (!data.value?.options) return false;

  return data.value.options.some((el) => {
    const contained = voteStore.findVoteByOptionId(el.id);
    // If contained is an array, check if it has any elements
    return Array.isArray(contained) ? contained.length > 0 : !!contained;
  });
});

const { data, error, status, refresh } = await useFetch(
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
