import {z} from "zod";
import {apiUrl} from "~/api/config";
import {User} from "./login.post";

const registerSchema = z.object({
  firstName: z.string(),
  lastName: z.string(),
  email: z.string().email(),
  password: z.string(),
});

type RegisterResponse = {
  user: User;
  access_token_expires_at: string;
  access_token: string;
  refresh_token: string;
  refresh_token_expires_at: string;
};

export default defineEventHandler(async (event) => {
  const body = await readValidatedBody(event, registerSchema.safeParse);

  if (!body.success) {
    throw body.error.issues;
  }

  const {firstName, lastName, email, password} = body.data;

  const response = await fetch(`${apiUrl}/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      last_name: lastName,
      first_name: firstName,
      email,
      password,
    }),
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.message || response.statusText);
  }

  return data as RegisterResponse;
});
