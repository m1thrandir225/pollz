import { apiUrl } from "~/api/config";
import type { PollWithOptions } from "~/types/poll";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");

  const response = await fetch(`${apiUrl}/polls/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    throw createError({
      status: response.status,
      statusMessage: response.statusText,
      message: response.statusText,
    });
  }

  const data = await response.json();

  return data as PollWithOptions;
});
