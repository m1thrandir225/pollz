<template>
  <form class="flex w-full flex-col items-start gap-4" @submit="onSubmit">
    <div
      v-for="option in options"
      :key="option.id"
      class="flex w-full flex-row items-center justify-between rounded-xl bg-neutral-100 p-4 transition-all duration-300 ease-in-out hover:bg-neutral-200 has-checked:bg-neutral-300 dark:bg-neutral-900 dark:hover:bg-neutral-800 dark:has-checked:bg-neutral-600"
      @click="optionId = option.id"
    >
      <label for="optionId"> {{ option.option_text }}</label>
      <input
        v-bind="optionIdAttrs"
        :id="'option-' + option.id"
        v-model="optionId"
        :value="option.id"
        type="radio"
        class="accent-orange-600"
      />
    </div>
    <button
      type="submit"
      class="mx-auto w-1/3 border border-transparent bg-orange-600 px-8 py-2 text-neutral-900 transition-all duration-100 ease-in-out hover:border-neutral-900 dark:text-neutral-100 dark:hover:border-neutral-100"
    >
      Vote
    </button>
    <p v-if="errors.optionId">
      {{ errors.optionId }}
    </p>
  </form>
</template>

<script setup lang="ts">
import * as z from "zod";

import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import type { PollOption } from "~/types/poll-option";

const voteStore = useVoteStore();

const props = defineProps<{
  options: PollOption[];
  refreshPageData: () => void;
}>();
const schema = z.object({
  optionId: z.string(),
});

const validationSchema = toTypedSchema(schema);

const { handleSubmit, defineField, errors } = useForm({
  validationSchema: validationSchema,
});

const [optionId, optionIdAttrs] = defineField("optionId");

const onSubmit = handleSubmit(async (values) => {
  const response = await $fetch("/api/vote", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(values),
  });

  voteStore.addVote(response);

  props.refreshPageData();
});
</script>
