import { apiUrl } from "~/api/config";
import type { RefreshTokenResponse } from "~/server/api/auth/refresh-token.post";

type RefreshTokenProps = {
  refreshToken: string;
};

export default async function (
  props: RefreshTokenProps,
): Promise<RefreshTokenResponse | null> {
  try {
    const response = await fetch(`${apiUrl}/refresh-token`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        refresh_token: props.refreshToken,
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error("please re-login");
    }

    return data as RefreshTokenResponse;
  } catch {
    return null;
  }
}
