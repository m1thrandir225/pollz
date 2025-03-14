<template>
  <h1>Editing Poll ID: {{ route.params.id }}</h1>
</template>

<script setup lang="ts">
definePageMeta({
  title: "Edit Poll",
  middleware: "auth",
});

const route = useRoute();
</script>
