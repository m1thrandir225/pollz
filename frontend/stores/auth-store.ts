import type { LoginResponse } from "~/server/api/auth/login.post";
import type { RefreshTokenResponse } from "~/server/api/auth/refresh-token.post";
import type { User } from "~/types/user";

export const useAuthStore = defineStore(
  "auth",
  () => {
    const accessToken = ref<string | null>(null);
    const accessTokenExpiresAt = ref<Date | null>(null);
    const refreshToken = ref<string | null>(null);
    const refreshTokenExpiresAt = ref<Date | null>(null);
    const user = ref<User | null>(null);
    const isRefreshing = ref(false);

    // Add a buffer time (e.g., 5 minutes) to refresh before expiration
    const TOKEN_REFRESH_BUFFER = 5 * 60 * 1000; // 5 minutes in milliseconds

    const shouldRefreshToken = computed(() => {
      const now = new Date();

      // First check if we have all the necessary data
      if (!accessToken.value || !refreshToken.value) {
        return false;
      }

      if (!accessTokenExpiresAt.value || !refreshTokenExpiresAt.value) {
        return false;
      }

      // Ensure we have proper Date objects
      const accessExpiry =
        accessTokenExpiresAt.value instanceof Date
          ? accessTokenExpiresAt.value
          : new Date(accessTokenExpiresAt.value);

      const refreshExpiry =
        refreshTokenExpiresAt.value instanceof Date
          ? refreshTokenExpiresAt.value
          : new Date(refreshTokenExpiresAt.value);

      // Check if refresh token is still valid
      if (now.getTime() >= refreshExpiry.getTime()) {
        return false; // Refresh token expired, can't refresh
      }

      // Check if access token needs refreshing (expired or about to expire)
      return now.getTime() >= accessExpiry.getTime() - TOKEN_REFRESH_BUFFER;
    });

    const isAuthenticated = computed(() => {
      const now = new Date();

      // Check if we have the necessary tokens
      if (!accessToken.value || !accessTokenExpiresAt.value) {
        return false;
      }

      // Ensure we have a proper Date object
      const accessExpiry =
        accessTokenExpiresAt.value instanceof Date
          ? accessTokenExpiresAt.value
          : new Date(accessTokenExpiresAt.value);

      // If we have a valid access token
      if (now.getTime() < accessExpiry.getTime()) {
        return true;
      }

      // If access token expired, check if refresh token is valid
      if (!refreshToken.value || !refreshTokenExpiresAt.value) {
        return false;
      }

      const refreshExpiry =
        refreshTokenExpiresAt.value instanceof Date
          ? refreshTokenExpiresAt.value
          : new Date(refreshTokenExpiresAt.value);

      return now.getTime() < refreshExpiry.getTime();
    });

    async function checkAndRefreshTokens(): Promise<boolean> {
      if (shouldRefreshToken.value && !isRefreshing.value) {
        return await refreshTokens();
      }
      return isAuthenticated.value;
    }

    async function refreshTokens(): Promise<boolean> {
      if (isRefreshing.value) {
        return isAuthenticated.value;
      }

      try {
        isRefreshing.value = true;

        if (!refreshToken.value || !user.value?.id) {
          logout();
          return false;
        }

        const refreshResponse = await $fetch("/api/auth/refresh-token", {
          method: "POST",
          body: JSON.stringify({
            refresh_token: refreshToken.value,
            user_id: user.value.id,
          }),
        });

        if (!refreshResponse) {
          logout();
          return false;
        }

        refreshAccess(refreshResponse);
        return true;
      } catch (error) {
        console.error("Token refresh failed:", error);
        logout();
        return false;
      } finally {
        isRefreshing.value = false;
      }
    }

    function login(response: LoginResponse) {
      accessToken.value = response.access_token;
      // Ensure we're storing Date objects
      accessTokenExpiresAt.value = new Date(response.access_token_expires_at);
      refreshToken.value = response.refresh_token;
      refreshTokenExpiresAt.value = new Date(response.refresh_token_expires_at);
      user.value = response.user;
    }

    function logout() {
      accessToken.value = null;
      accessTokenExpiresAt.value = null;
      refreshToken.value = null;
      refreshTokenExpiresAt.value = null;
      user.value = null;
    }

    function refreshAccess(response: RefreshTokenResponse) {
      accessToken.value = response.access_token;
      accessTokenExpiresAt.value = new Date(response.access_token_expires_at);
    }

    return {
      accessToken,
      accessTokenExpiresAt,
      refreshToken,
      refreshTokenExpiresAt,
      user,
      isAuthenticated,
      checkAndRefreshTokens,
      refreshTokens,
      login,
      logout,
      refreshAccess,
    };
  },
  {
    persist: {
      storage: piniaPluginPersistedstate.localStorage(),
    },
  },
);
