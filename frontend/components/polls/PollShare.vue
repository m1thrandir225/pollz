<template>
  <div class="relative flex w-full flex-col items-center gap-8">
    <div
      class="relative h-[200px] [&>svg]:mx-auto [&>svg]:h-[200px]"
      v-html="svgImage"
    ></div>
    <transition name="fade">
      <p v-if="showCopiedMessage">Copied to clipboard</p>
    </transition>
    <button
      type="button"
      class="cursor-pointer rounded-md bg-neutral-300 px-8 py-2 transition-all duration-300 ease-in-out hover:bg-neutral-200 dark:bg-neutral-600 dark:hover:bg-neutral-700"
      @click="copyUrl"
    >
      {{ route.fullPath }}
    </button>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  url: string;
}>();

const route = useRoute();

const showCopiedMessage = ref(false);

const svgImage = ref(createSVG(props.url));

function copyUrl() {
  if (navigator == null) return;

  navigator.clipboard.writeText(props.url);

  showCopiedMessage.value = true;

  setTimeout(() => {
    showCopiedMessage.value = false;
  }, 2000);
}
</script>
