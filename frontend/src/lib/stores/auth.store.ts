import { writable} from "svelte/store";
import { browser } from "$app/environment";

export const isAuthenticated = writable(false);
export const user = writable(null);

export async function checkAuth() {
    
}