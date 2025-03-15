export default defineNuxtRouteMiddleware((to, from) => {
  const accessToken = useCookie("access");
  const refreshToken = useCookie("refresh");
  const user = useCookie("user");

  const isAuthenticated = computed(
    () => accessToken.value && refreshToken.value
  );

  if (
    !isAuthenticated.value &&
    to.path !== "/login" &&
    to.path !== "/register"
  ) {
    return navigateTo("/login");
  } else if (
    isAuthenticated.value &&
    (to.path === "/login" || to.path === "/register")
  ) {
    return navigateTo("/");
  }
});
