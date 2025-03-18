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

  const data = await response.json();
  if (!response.ok) {
    throw createError({
      status: response.status,
      statusMessage: response.statusText,
      message: data.error,
    });
  }

  return data as PollWithOptions;
});
