const runtimeConfig = useRuntimeConfig();

export const baseUrl = runtimeConfig.public.apiSource;

export const apiUrl = `${baseUrl}/api/v1`;
