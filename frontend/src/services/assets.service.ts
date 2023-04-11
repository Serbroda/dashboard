import {BASE_URL} from "../config";

export async function fetchAsset(file: string): Promise<Response> {
    return fetch(`${BASE_URL}/api/v1/assets/${file}`)
}
