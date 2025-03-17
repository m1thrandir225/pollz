<template>
  <div
    class="flex flex-col items-center justify-center w-full min-h-screen gap-8"
  >
    <NuxtLink
      v-if="!authStore.isAuthenticated"
      to="/login"
      class="px-8 py-2 italic font-light text-white transition-all duration-100 ease-in-out bg-black border-2 rounded-sm hover:text-orange-600 hover:font-normal dark:bg-white dark:text-black hover:border-orange-400"
      >Get Started</NuxtLink
    >
    <NuxtLink
      v-else
      to="/create-poll"
      class="px-8 py-2 italic font-light text-white transition-all duration-100 ease-in-out bg-black border-2 rounded-sm hover:text-orange-600 hover:font-normal dark:bg-white dark:text-black hover:border-orange-400"
    >
      Create a poll
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">
const authStore = useAuthStore();

useSeoMeta({
  title: "Pollz",
});

definePageMeta({
  middleware: "auth",
});
</script>
