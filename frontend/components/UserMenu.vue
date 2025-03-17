<template>
  <div class="flex md:flex-row flex-col gap-2 items-center justify-start">
    <h1 class="font-mono dark:text-neutral-100 text-neutral-900">
      Hi, {{ user.first_name }} {{ user.last_name }}
    </h1>
  </div>
</template>

<script setup lang="ts">
import type { User } from "~/server/api/auth/login.post";

defineProps<{
  user: User;
}>();
</script>
