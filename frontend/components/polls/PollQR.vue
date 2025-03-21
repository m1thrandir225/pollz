<template>
  <div
    v-motion
    :initial="initial"
    :enter="enter"
    :leave="exit"
    class="relative h-[200px] w-full"
  >
    <div
      class="relative h-auto [&>svg]:mx-auto [&>svg]:h-[200px]"
      v-html="svgImage"
    ></div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  url: string;
}>();

const initial = ref({
  y: -100,
  opacity: 0,
});

const enter = ref({
  y: 0,
  opacity: 1,
  transition: {
    type: "spring",
    stiffness: 250,
    damping: 25,
    mass: 0.5,
  },
});

const exit = ref({
  y: -100,
  opacity: 0,
  transition: {
    type: "spring",
    stiffness: 250,
    damping: 25,
    mass: 0.5,
  },
});

const svgImage = ref(createSVG(props.url));
</script>
