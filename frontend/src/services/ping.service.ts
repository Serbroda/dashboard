import {BASE_URL} from "../config";

const ping = (target: string): Promise<Response> => {
    return fetch(`${BASE_URL}/api/v1/ping?target=${encodeURIComponent(target)}`);
}

export default ping;
