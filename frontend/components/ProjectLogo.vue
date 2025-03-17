<template>
  <NuxtLink
    to="/"
    :style="{ fontSize: size + 'px' }"
    class="w-auto font-mono font-bold text-center transition-all duration-150 ease-in-out group hover:text-orange-600 text-neutral-900 dark:text-neutral-100"
  >
    Po<span
      class="text-orange-600 transition-all duration-150 ease-in-out group-hover:text-neutral-900 dark:group-hover:text-neutral-100"
      >ll</span
    >z
  </NuxtLink>
</template>

<script setup lang="ts">
const { size = 96 } = defineProps<{
  size?: number;
}>();
</script>
