<template>
  <NuxtLink
    v-if="!authStore.isAuthenticated"
    to="/login"
    class="rounded-sm border-2 bg-black px-8 py-2 font-light text-white italic transition-all duration-100 ease-in-out hover:border-orange-400 hover:font-normal hover:text-orange-600 dark:bg-white dark:text-black"
    >Get Started</NuxtLink
  >
  <NuxtLink
    v-else
    to="/create-poll"
    class="rounded-sm border-2 bg-black px-8 py-2 font-light text-white italic transition-all duration-100 ease-in-out hover:border-orange-400 hover:font-normal hover:text-orange-600 dark:bg-white dark:text-black"
  >
    Create a poll
  </NuxtLink>
</template>

<script setup lang="ts">
const authStore = useAuthStore();
</script>
