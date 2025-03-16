<template>
  <div
    class="flex flex-col items-center justify-center w-full min-h-screen gap-8"
  >
    <div class="flex flex-col items-start gap-4">
      <NuxtLink
        to="/"
        class="w-full font-mono font-bold text-center transition-all duration-150 ease-in-out group hover:text-orange-600 text-8xl text-neutral-900 dark:text-neutral-100"
      >
        Po<span
          class="text-orange-600 transition-all duration-150 ease-in-out group-hover:text-neutral-900 hover:dark:text-neutral-100"
          >ll</span
        >z
      </NuxtLink>
      <p class="font-mono text-xl text-black dark:text-white">
        Easy fast polls with real-time updates
      </p>
    </div>
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
const authStore = useAuthStoreStore();

definePageMeta({
  middleware: "auth",
});
</script>
