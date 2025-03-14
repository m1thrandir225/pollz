<template>
  <div class="w-full flex flex-col gap-8 items-start">
    <h1 class="text-4xl font-bold font-mono">Login</h1>
    <form
      class="flex flex-col gap-4 items-center bg-neutral-100 p-8 rounded-md"
      @submit="onSubmit"
    >
      <div class="flex flex-col gap-2 items-start">
        <label for="email">Email</label>
        <input
          id="email"
          type="email"
          name="email"
          v-model="email"
          v-bind="emailAttrs"
          class="w-full bg-neutral-50 placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
        />
        <p v-if="errors.email" class="text-red-500">{{ errors.email }}</p>
      </div>
      <div class="flex flex-col gap-2 items-start">
        <label for="password">Password</label>
        <input
          id="password"
          v-model="password"
          type="password"
          name="password"
          v-bind="passwordAttrs"
          class="w-full bg-neutral-50 placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
        />
        <p v-if="errors.password" class="text-red-500">{{ errors.password }}</p>
      </div>
      <button
        type="submit"
        class="px-6 py-2 bg-blue-200 rounded-md text-blue-600 font-semibold hover:scale-[1.02] transition-transform ease-in-out duration-100"
      >
        Continue
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import * as zod from "zod";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";

const schema = zod.object({
  email: zod.string().email(),
  password: zod.string().min(6),
});

const validationSchema = toTypedSchema(schema);
const { handleSubmit, defineField, errors } = useForm({
  validationSchema: validationSchema,
});

const [email, emailAttrs] = defineField("email");
const [password, passwordAttrs] = defineField("password");

const onSubmit = handleSubmit(async (values) => {
  const response = await $fetch("/api/auth/login", {
    method: "POST",
    body: JSON.stringify(values),
  });

  console.log(response);
});
</script>
