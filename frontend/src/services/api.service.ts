import type {Config} from "../models/config.model";

const baseUrl = import.meta.env.VITE_API_BASE_URL || "";

class ApiService {
    constructor(private baseUrl: string) {
    }

    async getConfig(): Promise<Config> {
        const body = await fetch(`${this.baseUrl}/api/v1/config`);
        return body.json();
    }

    async getCustomCss(): Promise<Response> {
        return this.fetchStaticFile('custom.css');
    }

    async fetchStaticFile(file: string): Promise<Response> {
        return fetch(`${this.baseUrl}/static/${file}`)
    }
}

const apiService = new ApiService(baseUrl)

export {ApiService, apiService, baseUrl}
