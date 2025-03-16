export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStoreStore();
  if (
    !authStore.isAuthenticated &&
    to.path !== "/login" &&
    to.path !== "/register"
  ) {
    return navigateTo("/login");
  } else if (
    !authStore.isAuthenticated &&
    (to.path === "/login" || to.path === "/register")
  ) {
    return navigateTo("/");
  }
});
