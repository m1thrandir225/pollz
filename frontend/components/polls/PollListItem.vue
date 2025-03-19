<template>
  <NuxtLink
    :to="'/polls/' + item.id"
    class="group flex w-full flex-row items-center justify-between rounded-xl bg-neutral-100 px-4 py-4 transition-all duration-300 ease-in-out hover:bg-neutral-200 dark:bg-neutral-900 dark:hover:bg-neutral-800"
  >
    <h1
      class="font-mono text-lg text-neutral-900 group-hover:text-orange-600 dark:text-neutral-100"
    >
      {{ item.description }}
    </h1>
    <div class="flex flex-row items-center gap-4">
      <p class="font-mono text-[12px] text-neutral-900 dark:text-neutral-100">
        {{ formatedDate }}
      </p>
      <Icon
        name="material-symbols:arrow-right-alt"
        size="16px"
        mode="svg"
        class="text-neutral-900 duration-100 ease-in-out group-hover:scale-150 group-hover:text-orange-600 dark:text-neutral-100"
      />
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
import type { Poll } from "~/types/poll";

const props = defineProps<{
  item: Poll;
}>();

const formatedDate = ref(formatDate(new Date(props.item.active_until)));
</script>
