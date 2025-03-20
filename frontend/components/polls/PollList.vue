<template>
  <div
    class="mx-auto flex w-full max-w-[90%] flex-col items-start gap-8 lg:max-w-[650px]"
  >
    <div class="flex w-full flex-row justify-between">
      <h1 class="font-mono text-neutral-900 dark:text-neutral-100">
        Showing:
        <span class="text-[12px] text-orange-600">
          {{ showActive ? "current" : "expired" }} polls</span
        >
      </h1>
      <div class="flex flex-row gap-2">
        <button
          class="flex items-center justify-center rounded-md border border-neutral-400 px-4 py-2 text-neutral-900 transition-all duration-300 ease-in-out hover:border-transparent hover:bg-neutral-200 hover:text-orange-600 dark:text-neutral-100 dark:hover:bg-neutral-800"
          @click="toggleActive()"
        >
          <Icon
            :name="
              showActive
                ? 'material-symbols:clock-arrow-up'
                : 'material-symbols:clock-arrow-down'
            "
            size="18px"
            mode="css"
          />
        </button>
        <NuxtLink
          to="/create-poll"
          class="flex items-center justify-center rounded-md border border-neutral-400 px-4 py-2 text-neutral-900 transition-all duration-300 ease-in-out hover:border-transparent hover:bg-neutral-200 hover:text-orange-600 dark:text-neutral-100 dark:hover:bg-neutral-800"
        >
          <Icon name="material-symbols:add" size="18px" mode="css" />
        </NuxtLink>
      </div>
    </div>

    <div
      class="flex h-[600px] w-full flex-col items-start gap-4 overflow-y-scroll px-4 py-4"
    >
      <PollListItem
        v-for="item in showActive ? activePolls : inactivePolls"
        :key="item.id"
        :item="item"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Poll } from "~/types/poll";
import PollListItem from "./PollListItem.vue";

const showActive = ref(true);
const toggleActive = useToggle(showActive);

const props = defineProps<{
  list: Poll[];
}>();

const now = new Date();

const isPollActive = (activeUntil: string) => {
  const activeTime = new Date(activeUntil).getTime();
  return now.getTime() < activeTime;
};

const activePolls = computed(() => {
  return props.list.filter((item) => isPollActive(item.active_until));
});

const inactivePolls = computed(() => {
  return props.list.filter((item) => !isPollActive(item.active_until));
});
</script>
