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

    const TOKEN_REFRESH_BUFFER = 5 * 60 * 1000;

    const shouldRefreshToken = computed(() => {
      const now = new Date();

      if (!accessToken.value || !refreshToken.value) {
        return false;
      }

      if (!accessTokenExpiresAt.value || !refreshTokenExpiresAt.value) {
        return false;
      }

      const accessExpiry =
        accessTokenExpiresAt.value instanceof Date
          ? accessTokenExpiresAt.value
          : new Date(accessTokenExpiresAt.value);

      const refreshExpiry =
        refreshTokenExpiresAt.value instanceof Date
          ? refreshTokenExpiresAt.value
          : new Date(refreshTokenExpiresAt.value);

      if (now.getTime() >= refreshExpiry.getTime()) {
        return false;
      }

      return now.getTime() >= accessExpiry.getTime() - TOKEN_REFRESH_BUFFER;
    });

    const isAuthenticated = computed(() => {
      const now = new Date();

      if (!accessToken.value || !accessTokenExpiresAt.value) {
        return false;
      }

      const accessExpiry =
        accessTokenExpiresAt.value instanceof Date
          ? accessTokenExpiresAt.value
          : new Date(accessTokenExpiresAt.value);

      if (now.getTime() < accessExpiry.getTime()) {
        return true;
      }

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
