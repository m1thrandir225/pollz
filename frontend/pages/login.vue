<template>
  <form @submit="onSubmit">
    <input name="email" v-model="email" type="email" />
    <span v-if="errors.email">{{ errors.email }}</span>
    <input name="password" v-model="password" type="password" />
    <span>{{ errors.password }}</span>
    <button>Submit</button>
  </form>
</template>

<script setup lang="ts">
import {toTypedSchema} from "@vee-validate/zod";
import {useForm, useField} from "vee-validate";
import * as zod from "zod";

const schema = zod.object({
  email: zod.string().email(),
  password: zod.string().min(8),
});

const fieldSchema = toTypedSchema(schema);

const {handleSubmit, errors} = useForm({
  validationSchema: fieldSchema,
});

const {value: email} = useField("email");
const {value: password} = useField("password");

const onSubmit = handleSubmit((values) => {
  console.log(values);
});
</script>
