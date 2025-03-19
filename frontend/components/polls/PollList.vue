<template>
  <div
    class="flex flex-col items-start gap-8 lg:max-w-[650px] max-w-[90%] mx-auto w-full"
  >
    <div class="flex flex-row justify-between w-full">
      <h1 class="font-mono text-neutral-900 dark:text-neutral-100">
        Showing:
        <span class="text-orange-600 text-[12px]">
          {{ showActive ? "current" : "expired" }} polls</span
        >
      </h1>
      <div class="flex flex-row gap-2">
        <button
          class="px-4 py-2 border border-neutral-400 flex items-center justify-center rounded-md hover:bg-neutral-200 dark:hover:bg-neutral-800 hover:border-transparent hover:text-orange-600 transition-all ease-in-out duration-300 text-neutral-900 dark:text-neutral-100"
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
          class="px-4 py-2 border border-neutral-400 flex items-center justify-center rounded-md hover:bg-neutral-200 dark:hover:bg-neutral-800 hover:border-transparent hover:text-orange-600 transition-all ease-in-out duration-300 text-neutral-900 dark:text-neutral-100"
        >
          <Icon name="material-symbols:add" size="18px" mode="css" />
        </NuxtLink>
      </div>
    </div>

    <div
      class="flex flex-col items-start w-full gap-4 h-[600px] overflow-y-scroll px-4 py-4"
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
