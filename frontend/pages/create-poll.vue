<template>
  <div class="mx-auto flex max-w-lg items-center justify-center">
    <form class="flex w-full flex-col gap-8" @submit="onSubmit">
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
          class="sborder-slate-200 ease w-full min-w-[250px] border bg-neutral-900 px-3 py-2 text-sm text-neutral-100 shadow-sm transition duration-300 placeholder:font-mono placeholder:text-neutral-400 hover:border-orange-300 focus:border-orange-400 focus:shadow focus:outline-none dark:bg-neutral-50 dark:text-neutral-900 dark:placeholder:text-neutral-400"
        />
        <p v-if="errors.description" class="max-w-[150px] text-sm text-red-500">
          {{ errors.description }}
        </p>
      </div>
      <div class="flex flex-row items-center justify-between gap-2">
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
          class="font-mono text-neutral-900 dark:text-neutral-100"
          >Options</label
        >
        <div class="flex w-full flex-row items-center justify-between gap-4">
          <input
            id="options"
            v-model="optionInput"
            type="text"
            name="options"
            placeholder="add a poll option"
            class="ease w-full min-w-[250px] border border-slate-200 bg-neutral-900 px-3 py-2 text-sm text-neutral-100 shadow-sm transition duration-300 placeholder:font-mono placeholder:text-neutral-400 hover:border-orange-300 focus:border-orange-400 focus:shadow focus:outline-none dark:bg-neutral-50 dark:text-neutral-900 dark:placeholder:text-neutral-400"
          />
          <button type="button" class="w-auto" @click="addOption()">Add</button>
        </div>
        <p v-if="errors.description" class="max-w-[150px] text-sm text-red-500">
          {{ errors.description }}
        </p>
      </div>
      <div
        v-if="options && options?.length > 0"
        class="flex max-h-[300px] w-full flex-col gap-4 overflow-y-scroll border border-neutral-900 p-4 dark:border-neutral-100"
      >
        <div
          v-for="(option, index) in options"
          :key="index"
          class="dark:border-neural-600 flex w-full flex-row items-center justify-between border-b border-neutral-200 px-4 py-4"
        >
          {{ option }}
          <button
            class="flex cursor-pointer items-center justify-center transition-all duration-300 ease-in-out hover:text-red-500"
            @click="removeOption(index)"
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
        class="w-full border border-transparent bg-orange-600 px-8 py-2 text-neutral-900 transition-all duration-100 ease-in-out hover:border-neutral-900 dark:text-neutral-100 dark:hover:border-neutral-100"
      >
        Continue
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

definePageMeta({
  middleware: "auth",
});

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
const [options] = defineField("options");

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
        Authorization: `Bearer ${authStore.accessToken!}`,
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
