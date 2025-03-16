import type { RefreshTokenResponse } from "~/server/api/auth/refresh-token.post";

export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useAuthStoreStore();

  if (!authStore.isAuthenticated && to.path !== "/login") {
    return navigateTo("/login");
  }

  if (authStore.isAuthenticated && to.path === "/login") {
    return navigateTo("/");
  }

  if (authStore.isAuthenticated && to.path === "/register") {
    return navigateTo("/");
  }

  if (authStore.isAuthenticated && authStore.shouldRefreshToken) {
    try {
      const response = await $fetch("/api/auth/refresh-token", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${authStore.accessToken}`,
        },
        body: JSON.stringify({
          refresh_token: authStore.refreshToken,
        }),
      });

      authStore.refreshAccess(response);
    } catch (error) {
      authStore.logout();
      return navigateTo("/login");
    }
  }
});
