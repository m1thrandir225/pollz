export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useAuthStore();

  const isAuth = await authStore.checkAndRefreshTokens();
  if (!isAuth && to.path !== "/login" && to.path !== "/register") {
    return navigateTo("/login");
  }

  if (isAuth && ["/login", "/register"].includes(to.path)) {
    return navigateTo("/");
  }
});
