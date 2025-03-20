<template>
  <div class="my-8 flex w-full flex-col items-center gap-8 overflow-hidden">
    <PollList v-if="data" :list="data" />
  </div>
</template>
<script setup lang="ts">
import PollList from "~/components/polls/PollList.vue";

const authStore = useAuthStore();

const { data, refresh, clear, error, status } = useAsyncData("user-polls", () =>
  $fetch("/api/polls", {
    method: "GET",
    headers: {
      Authorization: `Bearer ${authStore.accessToken}`,
    },
  }),
);

if (error.value) {
  throw error;
}

useSeoMeta({
  title: "Pollz",
});

definePageMeta({
  middleware: "auth",
});
</script>
