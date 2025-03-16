import type { LoginResponse, User } from "~/server/api/auth/login.post";

export const useAuthStoreStore = defineStore("auth", () => {
  const accessToken = useCookie("access");
  const refreshToken = useCookie("refresh");
  const user = useCookie("user");

  const isAuthenticated = computed(() => !!accessToken.value);
  const shouldRefresh = computed(() => {
    //No refresh token or no user
    if (!refreshToken.value || !user.value) return false;

    //Has access token
    if (accessToken.value) return false;

    return true;
  });

  function getUserDetails() {
    if (!user.value) return null;
    return JSON.parse(user.value) as User;
  }

  return { isAuthenticated, shouldRefresh, getUserDetails };
});
