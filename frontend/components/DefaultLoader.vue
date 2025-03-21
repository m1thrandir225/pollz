<template>
  <transition name="fade">
    <Icon
      v-if="show"
      name="codex:loader"
      size="24px"
      mode="svg"
      class="text-orange-600"
    />
  </transition>
</template>

<script setup lang="ts">
defineProps<{
  show: boolean;
}>();
</script>
