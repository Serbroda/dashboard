import type {Writable} from "svelte/store";
import {writable} from "svelte/store";

export type FetchStore<T> = [Writable<T>, Writable<boolean>, Writable<boolean>, () => Promise<void>];

const fetchStore = <T>(url: string, init?: RequestInit): FetchStore<T> => {
    const loading = writable<boolean>(false)
    const error = writable<boolean>(false)
    const data = writable<T>();

    async function get() {
        loading.set(true)
        error.set(false)
        try {
            const res = await fetch(url, init);
            data.set(await res.json());
        } catch (e) {
            error.set(e)
        } finally {
            loading.set(false)
        }
    }

    get()

    return [data, loading, error, get]
}

export default fetchStore;
