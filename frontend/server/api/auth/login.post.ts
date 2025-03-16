import { z } from "zod";
import { apiUrl } from "~/api/config";

const loginSchema = z.object({
  email: z.string().email(),
  password: z.string(),
});

export type User = {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  created_at: string;
};

export type LoginResponse = {
  user: User;
  access_token_expires_at: string;
  access_token: string;
  refresh_token: string;
  refresh_token_expires_at: string;
};

export default defineEventHandler(async (event) => {
  const body = await readValidatedBody(event, loginSchema.safeParse);

  if (!body.success) {
    throw body.error.issues;
  }

  const { email, password } = body.data;

  const response = await fetch(`${apiUrl}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.message || response.statusText);
  }

  return data as LoginResponse;
});
