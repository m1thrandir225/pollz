<template>
  <div class="flex flex-row items-center justify-between w-full">
    <ProjectLogo :size="36" class="w-32" />
    <div
      v-if="authStore.isAuthenticated && authStore.user"
      class="flex flex-row gap-4 items-center"
    >
      <UserMenu :user="authStore.user" />
      <button @click="authStore.logout">
        <Icon
          name="bx:log-out"
          class="text-neutral-900 dark:text-neutral-100"
          size="18px"
          mode="svg"
        />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
const authStore = useAuthStore();
</script>
