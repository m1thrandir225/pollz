<script setup lang="ts">
const isDark = useDark();

const toggleDark = useToggle(isDark);
</script>

<template>
  <button
    class="text-neutral-900 dark:text-neutral-100 hover:text-orange-600"
    @click="toggleDark()"
  >
    <Icon
      v-if="isDark"
      name="material-symbols:dark-mode-outline-rounded"
      size="24px"
      mode="svg"
    />
    <Icon
      v-else
      name="material-symbols:sunny-outline-rounded"
      size="24px"
      mode="svg"
    />
  </button>
</template>
