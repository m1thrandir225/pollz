<template>
  <div class="mx-auto h-auto max-w-[300px]">
    <Doughnut :data="chartData" :options="chartOptions" />
  </div>
</template>

<script setup lang="ts">
import { Doughnut } from "vue-chartjs";
import type { ChartData, ChartOptions } from "chart.js";
import { Chart as ChartJs, ArcElement, Tooltip, Legend } from "chart.js";

ChartJs.register(ArcElement, Tooltip, Legend);

const chartOptions = ref<ChartOptions<"doughnut">>({
  responsive: true,
  maintainAspectRatio: true,
});

defineProps<{
  chartData: ChartData<"doughnut", number[], unknown>;
}>();
</script>
