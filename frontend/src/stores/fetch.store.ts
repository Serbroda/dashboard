import type {Writable} from "svelte/store";
import {writable} from "svelte/store";

export type FetchStore = [Writable<Response>, Writable<boolean>, Writable<boolean>, () => Promise<void>];

const fetchStore = (url: string, init?: RequestInit): FetchStore => {
    const loading = writable<boolean>(false)
    const error = writable<boolean>(false)
    const response = writable<Response>();

    async function get() {
        loading.set(true)
        error.set(false)
        try {
            const res = await fetch(url, init);
            response.set(res);
        } catch (e) {
            error.set(e)
        } finally {
            loading.set(false)
        }
    }

    get()

    return [response, loading, error, get]
}

export default fetchStore;
