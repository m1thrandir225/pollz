import * as z from "zod";
import { apiUrl } from "~/api/config";
import type { PollOption } from "~/types/poll-option";

const schema = z.object({
  poll_id: z.string(),
  option_text: z.string(),
});

export default defineEventHandler(async (event) => {
  const body = await readValidatedBody(event, schema.safeParse);

  if (!body.success) {
    throw body.error.issues;
  }

  const authToken = event.headers.get("Authorization");

  if (!authToken) {
    throw createError({
      status: 403,
      statusMessage: "Missing bearer token",
      message: "Missing bearer token",
    });
  }
  const response = await fetch(`${apiUrl}/poll-options`, {
    method: "POST",
    headers: { Authorization: authToken, "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });

  const data = await response.json();

  if (!response.ok) {
    throw createError({
      status: response.status,
      statusMessage: response.statusText,
      message: data.error,
    });
  }
  return data as PollOption;
});
