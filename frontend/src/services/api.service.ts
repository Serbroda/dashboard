class ApiService {
    constructor(private baseUrl: string) {
    }

    async get(): Promise<any> {
        const body = await fetch(`${this.baseUrl}/api/v1/config`);
        return body.json();
    }
}

const baseUrl = import.meta.env.VITE_API_BASE_URL || "";
const apiService = new ApiService(baseUrl)

export {ApiService, apiService}
