<!-- eslint-disable vue/html-self-closing -->
<template>
  <div class="flex flex-col items-start w-full gap-8">
    <h1 class="font-mono text-4xl font-bold">Login</h1>
    <form
      class="flex flex-col items-center gap-4 p-8 rounded-md bg-neutral-100"
      @submit="onSubmit"
    >
      <div class="flex flex-col items-start gap-2">
        <label for="email">Email</label>
        <input
          id="email"
          v-model="email"
          type="email"
          name="email"
          v-bind="emailAttrs"
          class="w-full px-3 py-2 text-sm transition duration-300 border rounded-md shadow-sm bg-neutral-50 placeholder:text-slate-400 text-slate-700 border-slate-200 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 focus:shadow"
        />
        <p v-if="errors.email" class="text-red-500">{{ errors.email }}</p>
      </div>
      <div class="flex flex-col items-start gap-2">
        <label for="password">Password</label>
        <input
          id="password"
          v-model="password"
          type="password"
          name="password"
          v-bind="passwordAttrs"
          class="w-full px-3 py-2 text-sm transition duration-300 border rounded-md shadow-sm bg-neutral-50 placeholder:text-slate-400 text-slate-700 border-slate-200 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 focus:shadow"
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

const authStore = useAuthStoreStore();

definePageMeta({
  middleware: "auth",
});

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
  try {
    const response = await $fetch("/api/auth/login", {
      method: "POST",
      body: JSON.stringify(values),
    });

    authStore.login(response);
  } catch (error) {
    console.error(error);
  }
});
</script>
