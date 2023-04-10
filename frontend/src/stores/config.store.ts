import {writable} from "svelte/store";
import type {Config} from "../models/config.model";
import {apiService} from "../services/api.service";

// https://svelte.dev/repl/b2d671b8119845ca903667f1b3a96e31?version=3.37.0
export default function () {
    const loading = writable<boolean>(false)
    const error = writable<boolean>(false)
    const data = writable<Config>({sections: [], title: "hedywyd?y!"})

    async function get() {
        loading.set(true)
        error.set(false)
        try {
            const response = await apiService.getConfig();
            data.set(response);
        } catch(e) {
            error.set(e)
        } finally {
            loading.set(false)
        }
    }

    get()

    return [ data, loading, error, get]
}
