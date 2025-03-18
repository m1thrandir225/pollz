<script setup lang="ts">
import type { NuxtError } from "#app";

const props = defineProps({
  error: Object as () => NuxtError,
});

const handleError = () => clearError({ redirect: "/" });
</script>

<template>
  <div
    class="w-full min-h-screen flex items-center justify-center flex-col bg-neutral-100 dark:bg-neutral-900"
  >
    <ProjectLogo :size="64" />

    <h2 class="text-neutral-900 dark:text-neutral-100 font-mono">
      Status: {{ error?.statusCode }}
    </h2>

    <button
      class="text-mono text-neutral-900 dark:text-neutral-100 cursor-pointer"
      @click="handleError"
    >
      Go Home
    </button>
  </div>
</template>
