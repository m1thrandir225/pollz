<!-- eslint-disable vue/html-self-closing -->
<template>
  <div
    class="flex min-h-[650px] w-full flex-col items-center justify-center gap-8"
  >
    <h1
      class="font-mono text-4xl font-bold text-neutral-900 dark:text-neutral-100"
    >
      Login
    </h1>
    <form
      class="flex flex-col items-center gap-4 rounded-sm border border-neutral-900 p-16 dark:border-neutral-100"
      @submit="onSubmit"
    >
      <div class="flex flex-col items-start gap-2">
        <label
          for="email"
          class="font-mono text-neutral-900 dark:text-neutral-100"
          >Email</label
        >
        <input
          id="email"
          v-model="email"
          type="email"
          name="email"
          placeholder="your email"
          v-bind="emailAttrs"
          class="ease w-full border border-slate-200 bg-neutral-900 px-3 py-2 text-sm text-neutral-100 shadow-sm transition duration-300 placeholder:font-mono placeholder:text-neutral-400 hover:border-orange-300 focus:border-orange-400 focus:shadow focus:outline-none dark:bg-neutral-50 dark:text-neutral-900 dark:placeholder:text-neutral-400"
        />
        <p v-if="errors.email" class="max-w-[150px] text-sm text-red-500">
          {{ errors.email }}
        </p>
      </div>
      <div class="flex flex-col items-start gap-2">
        <label
          for="password"
          class="font-mono text-neutral-900 dark:text-neutral-100"
          >Password</label
        >
        <input
          id="password"
          v-model="password"
          type="password"
          placeholder="your password"
          name="password"
          v-bind="passwordAttrs"
          class="ease w-full border border-slate-200 bg-neutral-900 px-3 py-2 text-sm text-neutral-100 shadow-sm transition duration-300 placeholder:font-mono placeholder:text-neutral-400 hover:border-orange-300 focus:border-orange-400 focus:shadow focus:outline-none dark:bg-neutral-50 dark:text-neutral-900 dark:placeholder:text-neutral-400"
        />
        <p v-if="errors.password" class="max-w-[150px] text-sm text-red-500">
          {{ errors.password }}
        </p>
      </div>
      <button
        type="submit"
        class="w-full border border-transparent bg-orange-600 px-8 py-2 text-neutral-900 transition-all duration-100 ease-in-out hover:border-neutral-900 dark:text-neutral-100 dark:hover:border-neutral-100"
      >
        Continue
      </button>
    </form>
    <p class="font-mono text-neutral-900 dark:text-neutral-100">
      Don't have an account?
      <NuxtLink to="/register" class="font-bold text-orange-600"
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
