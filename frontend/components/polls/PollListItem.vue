<template>
  <NuxtLink
    :to="'/polls/' + item.id"
    class="group flex flex-row items-center justify-between w-full bg-neutral-100 dark:bg-neutral-900 hover:bg-neutral-200 dark:hover:bg-neutral-800 transition-all ease-in-out duration-300 rounded-xl px-4 py-4"
  >
    <h1
      class="text-lg font-mono group-hover:text-orange-600 text-neutral-900 dark:text-neutral-100"
    >
      {{ item.description }}
    </h1>
    <div class="flex flex-row items-center gap-4">
      <p class="text-[12px] font-mono text-neutral-900 dark:text-neutral-100">
        {{ formatedDate }}
      </p>
      <Icon
        name="material-symbols:arrow-right-alt"
        size="16px"
        mode="svg"
        class="group-hover:scale-150 group-hover:text-orange-600 ease-in-out duration-100 text-neutral-900 dark:text-neutral-100"
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
