import type { LoginResponse, User } from "~/server/api/auth/login.post";
import type { RefreshTokenResponse } from "~/server/api/auth/refresh-token.post";

export const useAuthStoreStore = defineStore(
  "auth",
  () => {
    const accessToken = ref<string | null>(null);
    const accessTokenExpiresAt = ref<Date | null>(null);

    const refreshToken = ref<string | null>(null);
    const refreshTokenExpiresAt = ref<Date | null>(null);

    const user = ref<User | null>(null);

    const shouldRefreshToken = computed(() => {
      if (!accessTokenExpiresAt.value || !refreshTokenExpiresAt.value) {
        return false;
      }

      if (!accessToken.value || !refreshToken.value) {
        return false;
      }

      return (
        new Date() > accessTokenExpiresAt.value &&
        new Date() < refreshTokenExpiresAt.value
      );
    });

    const isAuthenticated = computed(() => !!accessToken.value);

    function login(response: LoginResponse) {
      accessToken.value = response.access_token;
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
      shouldRefreshToken,
      isAuthenticated,
      login,
      logout,
      refreshAccess,
    };
  },
  {
    persist: {
      storage: piniaPluginPersistedstate.cookies({
        httpOnly: true,
      }),
    },
  }
);
