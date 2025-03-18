import { apiUrl } from "~/api/config";
import type { User } from "~/types/user";
export default defineEventHandler(async (event) => {
  const authToken = event.headers.get("Authorization");

  if (!authToken) {
    throw createError({
      status: 401,
      statusText: "Missing bearer token",
      message: "Missing bearer token",
    });
  }

  const response = await fetch(`${apiUrl}/me`, {
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

  return data as User;
});
