import { apiUrl } from "~/api/config";
import type { Poll } from "~/types/poll";

/**
 * API route for all the polls the user has (both active and inactive)
 */

export default defineEventHandler(async (event) => {
  const authToken = event.headers.get("Authorization");

  if (!authToken) {
    throw createError({
      status: 403,
      statusMessage: "Missing bearer token",
      message: "Missing bearer token",
    });
  }

  const response = await fetch(`${apiUrl}/polls`, {
    method: "GET",
    headers: {
      Authorization: authToken,
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

  return data as Poll[];
});
