<template>
  <div class="flex flex-col items-center justify-start gap-2 md:flex-row">
    <h1 class="text-[14px]">Hi, {{ user.first_name }} {{ user.last_name }}</h1>
  </div>
</template>

<script setup lang="ts">
import type { User } from "~/types/user";

defineProps<{
  user: User;
}>();
</script>
