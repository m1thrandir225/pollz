import * as z from "zod";
import { apiUrl } from "~/api/config";
import type { Vote } from "~/types/vote";

const schema = z.object({
  optionId: z.string(),
});

export default defineEventHandler(async (event) => {
  const body = await readValidatedBody(event, schema.safeParse);

  if (!body.success) {
    throw body.error.issues;
  }

  const ipAddress = getRequestIP(event, { xForwardedFor: true });
  const userAgent = event.headers.get("user-agent");

  const response = await fetch(`${apiUrl}/vote`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      ip_address: ipAddress,
      user_agent: userAgent,
      option_id: body.data.optionId,
    }),
  });

  const data = await response.json();

  if (!response.ok) {
    throw createError({
      status: response.status,
      statusMessage: response.statusText,
      message: data.error,
    });
  }

  return data as Vote;
});
