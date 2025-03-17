<!-- eslint-disable vue/html-self-closing -->
<template>
  <div
    class="flex flex-col items-center justify-center w-full gap-8 min-h-[800px]"
  >
    <h1
      class="font-mono text-4xl font-bold text-neutral-900 dark:text-neutral-100"
    >
      Login
    </h1>
    <form
      class="flex flex-col items-center gap-4 p-16 rounded-sm border border-neutral-900 dark:border-neutral-100"
      @submit="onSubmit"
    >
      <div class="flex flex-col items-start gap-2">
        <label
          for="email"
          class="text-neutral-900 dark:text-neutral-100 font-mono"
          >Email</label
        >
        <input
          id="email"
          v-model="email"
          type="email"
          name="email"
          placeholder="your email"
          v-bind="emailAttrs"
          class="w-full px-3 py-2 text-sm transition duration-300 border shadow-sm bg-neutral-900 dark:bg-neutral-50 placeholder:text-neutral-400 text-neutral-100 dark:text-neutral-900 dark:placeholder:text-neutral-400 border-slate-200 ease focus:outline-none focus:border-orange-400 hover:border-orange-300 focus:shadow placeholder:font-mono"
        />
        <p v-if="errors.email" class="text-red-500 max-w-[150px] text-sm">
          {{ errors.email }}
        </p>
      </div>
      <div class="flex flex-col items-start gap-2">
        <label
          for="password"
          class="text-neutral-900 dark:text-neutral-100 font-mono"
          >Password</label
        >
        <input
          id="password"
          v-model="password"
          type="password"
          placeholder="your password"
          name="password"
          v-bind="passwordAttrs"
          class="w-full px-3 py-2 text-sm transition duration-300 border shadow-sm bg-neutral-900 dark:bg-neutral-50 placeholder:text-neutral-400 text-neutral-100 dark:text-neutral-900 dark:placeholder:text-neutral-400 border-slate-200 ease focus:outline-none focus:border-orange-400 hover:border-orange-300 focus:shadow placeholder:font-mono"
        />
        <p v-if="errors.password" class="text-red-500 max-w-[150px] text-sm">
          {{ errors.password }}
        </p>
      </div>
      <button
        type="submit"
        class="w-full bg-orange-600 dark:text-neutral-100 text-neutral-900 px-8 py-2 rounded-md border border-transparent hover:border-neutral-900 dark:hover:border-neutral-100 transition-all ease-in-out duration-100"
      >
        Continue
      </button>
    </form>
    <p class="font-mono">
      Don't have an account?
      <NuxtLink to="/register" class="text-orange-600 font-bold"
        >Register</NuxtLink
      >
    </p>
  </div>
</template>

<script setup lang="ts">
import * as zod from "zod";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";

const authStore = useAuthStore();

useSeoMeta({
  title: "Login | Pollz",
});

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

    return navigateTo("/");
  } catch (error) {
    console.error(error);
  }
});
</script>
