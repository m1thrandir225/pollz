<template>
  <form class="max-w-lg mx-auto flex flex-col gap-8" @submit="onSubmit">
    <div class="flex flex-col items-start gap-2">
      <label
        for="description"
        class="text-neutral-900 dark:text-neutral-100 font-mono"
        >Description</label
      >
      <input
        id="description"
        v-model="description"
        type="text"
        name="description"
        placeholder="what's this poll about?"
        v-bind="descriptionAttrs"
        class="w-full min-w-[250px] px-3 py-2 text-sm transition duration-300 border shadow-sm bg-neutral-900 dark:bg-neutral-50 placeholder:text-neutral-400 text-neutral-100 dark:text-neutral-900 dark:placeholder:text-neutral-400 border-slate-200 ease focus:outline-none focus:border-orange-400 hover:border-orange-300 focus:shadow placeholder:font-mono"
      />
      <p v-if="errors.description" class="text-red-500 max-w-[150px] text-sm">
        {{ errors.description }}
      </p>
    </div>
    <div class="flex flex-row justify-between items-center gap-2">
      <label
        for="active_until"
        class="text-neutral-900 dark:text-neutral-100 font-mono"
        >Active Time</label
      >
      <select
        v-bind="activeUntilAttrs"
        id="active_until"
        v-model="activeUntil"
        name="active_until"
        class="text-neutral-900 dark:text-neutral-100 font-mono"
      >
        <option
          v-for="(activeTime, index) in selectActiveTimes"
          :key="index"
          class="text-neutral-900 dark:text-neutral-100 font-mono"
          :value="activeTime"
        >
          {{ activeTime }}
        </option>
      </select>
    </div>
    <button
      type="submit"
      class="w-full bg-orange-600 dark:text-neutral-100 text-neutral-900 px-8 py-2 border border-transparent hover:border-neutral-900 dark:hover:border-neutral-100 transition-all ease-in-out duration-100"
    >
      Continue
    </button>
  </form>
</template>

<script setup lang="ts">
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const selectActiveTimes = ref(["1 day", "2 day", "1 week"]);

const authStore = useAuthStore();

const schema = z.object({
  description: z.string().min(10),
  active_until: z
    .string()
    .refine((val) => selectActiveTimes.value.includes(val)),
});
const validationSchema = toTypedSchema(schema);
const { defineField, errors, handleSubmit } = useForm({
  validationSchema: validationSchema,
  initialValues: {
    active_until: selectActiveTimes.value[0],
  },
});

const [description, descriptionAttrs] = defineField("description");
const [activeUntil, activeUntilAttrs] = defineField("active_until");

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

    const userId = authStore.user?.id;

    if (!userId) throw new Error("user id missing");

    const result = await $fetch("/api/polls", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${authStore.accessToken!!}`,
      },
      body: JSON.stringify({
        description: values.description,
        active_until: utcDateValue,
        user_id: userId,
      }),
    });

    return navigateTo("/polls/" + result.id);
  } catch (error) {
    console.error(error);
  }
});
</script>
