export default defineNuxtRouteMiddleware((to, from) => {
  const cookie = useCookie("user");

  console.log("value: ", cookie.value);
  if (!cookie.value) {
    return navigateTo("/login");
  }
});
