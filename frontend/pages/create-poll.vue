<template>
  <div class="max-w-lg flex items-center mx-auto justify-center">
    <form class="w-full flex flex-col gap-8" @submit="onSubmit">
      <h1 class="font-mono text-xl font-bold">Create a poll</h1>
      <div class="flex flex-col items-start gap-2">
        <label for="description" class="font-mono">Description</label>
        <input
          id="description"
          v-model="description"
          type="text"
          name="description"
          placeholder="what's this poll about?"
          v-bind="descriptionAttrs"
          class="w-full min-w-[250px] px-3 py-2 text-sm transition duration-300 border shadow-sm bg-neutral-900 dark:bg-neutral-50 placeholder:text-neutral-400 dark:placeholder:text-neutral-400 text-neutral-100 dark:text-neutral-900 sborder-slate-200 ease focus:outline-none focus:border-orange-400 hover:border-orange-300 focus:shadow placeholder:font-mono"
        />
        <p v-if="errors.description" class="text-red-500 max-w-[150px] text-sm">
          {{ errors.description }}
        </p>
      </div>
      <div class="flex flex-row justify-between items-center gap-2">
        <label for="active_until" class="font-mono">Active Time</label>
        <select
          v-bind="activeUntilAttrs"
          id="active_until"
          v-model="activeUntil"
          name="active_until"
          class="font-mono"
        >
          <option
            v-for="(activeTime, index) in selectActiveTimes"
            :key="index"
            class="font-mono"
            :value="activeTime"
          >
            {{ activeTime }}
          </option>
        </select>
      </div>
      <div class="flex flex-col items-start gap-2">
        <label
          for="options"
          class="text-neutral-900 dark:text-neutral-100 font-mono"
          >Options</label
        >
        <div class="flex flex-row items-center justify-between gap-4 w-full">
          <input
            v-model="optionInput"
            id="options"
            type="text"
            name="options"
            placeholder="add a poll option"
            class="w-full min-w-[250px] px-3 py-2 text-sm transition duration-300 border shadow-sm bg-neutral-900 dark:bg-neutral-50 placeholder:text-neutral-400 text-neutral-100 dark:text-neutral-900 dark:placeholder:text-neutral-400 border-slate-200 ease focus:outline-none focus:border-orange-400 hover:border-orange-300 focus:shadow placeholder:font-mono"
          />
          <button type="button" class="w-auto" @click="addOption()">Add</button>
        </div>
        <p v-if="errors.description" class="text-red-500 max-w-[150px] text-sm">
          {{ errors.description }}
        </p>
      </div>
      <div
        class="flex flex-col w-full max-h-[300px] overflow-y-scroll p-4 border border-neutral-900 dark:border-neutral-100 gap-4"
        v-if="options && options?.length > 0"
      >
        <div
          v-for="(option, index) in options"
          :key="index"
          class="w-full flex flex-row items-center justify-between px-4 py-4 border-b border-neutral-200 dark:border-neural-600"
        >
          {{ option }}
          <button
            @click="removeOption(index)"
            class="cursor-pointer flex items-center justify-center hover:text-red-500 transition-all ease-in-out duration-300"
          >
            <Icon
              name="material-symbols:delete-outline"
              size="24px"
              mode="css"
            />
          </button>
        </div>
      </div>

      <button
        type="submit"
        class="w-full bg-orange-600 dark:text-neutral-100 text-neutral-900 px-8 py-2 border border-transparent hover:border-neutral-900 dark:hover:border-neutral-100 transition-all ease-in-out duration-100"
      >
        Continue
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { useField, useFieldArray, useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const selectActiveTimes = ref(["1 day", "2 day", "1 week"]);

const authStore = useAuthStore();

const optionInput = ref("");

const schema = z.object({
  description: z.string().min(2),
  active_until: z
    .string()
    .refine((val) => selectActiveTimes.value.includes(val)),
  options: z.array(z.string()).min(2),
});
const validationSchema = toTypedSchema(schema);
const { defineField, errors, handleSubmit } = useForm({
  validationSchema: validationSchema,
  initialValues: {
    active_until: selectActiveTimes.value[0],
    options: [],
  },
});

const [description, descriptionAttrs] = defineField("description");
const [activeUntil, activeUntilAttrs] = defineField("active_until");
const [options, optionsAttrs] = defineField("options");

const addOption = () => {
  if (optionInput.value.trim()) {
    const updatedOptions = [...(options.value || []), optionInput.value.trim()];
    options.value = updatedOptions;
    optionInput.value = "";
  }
};

const removeOption = (index: number) => {
  const updatedOptions = [...(options.value || [])];
  updatedOptions.splice(index, 1);
  options.value = updatedOptions;
};

const activeTimeToDate = ({
  selectedValue,
}: {
  selectedValue: string;
}): string => {
  const now = new Date();
  switch (selectedValue) {
    case "1 day": {
      return addDaysToDate(now, 1).toISOString();
    }
    case "2 day": {
      return addDaysToDate(now, 1).toISOString();
    }
    case "1 week": {
      return addWeeksToDate(now, 1).toISOString();
    }
    default:
      throw new Error("value not part of range");
  }
};

const onSubmit = handleSubmit(async (values) => {
  try {
    const utcDateValue = activeTimeToDate({
      selectedValue: values.active_until,
    });

    const result = await $fetch("/api/polls", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${authStore.accessToken!!}`,
      },
      body: JSON.stringify({
        description: values.description,
        active_until: utcDateValue,
        options: values.options,
      }),
    });

    return navigateTo("/polls/" + result.id);
  } catch (error) {
    console.error(error);
  }
});
</script>
