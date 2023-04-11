import type {Config} from "../models/config.model";
import fetchStore from "./fetch.store";
import {BASE_URL} from "../config";

// https://svelte.dev/repl/b2d671b8119845ca903667f1b3a96e31?version=3.37.0


export default function() {
    return fetchStore<Config>(`${BASE_URL}/api/v1/config`)
}

export function saveConfig(config: Config): Promise<Response> {
    return fetch(`${BASE_URL}/api/v1/config`, {
        method: "POST",
        body: JSON.stringify(config),
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    })
}