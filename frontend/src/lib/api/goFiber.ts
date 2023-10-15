import { variables } from '$lib/variables';
import axios from 'axios';

export class GoRestClient {
    baseUrl: string | undefined;
    debug: boolean = false
    constructor(debug: boolean = false, group?: string, override?: boolean) {
        if (!override) {
            this.baseUrl = `http://${variables.goBackendURL}:${variables.goBackendPort}`;
        } else {
            this.baseUrl = group;
        }
        this.debug = debug
        if (debug) {
            console.log(this.baseUrl, group);
        }
    }

    async get(endpoint: string, id?: string | number): Promise<any> {
        try {
            let res: any;
            if (id) {
                res = await axios.get(`${this.baseUrl}/${endpoint}/${id}`);
            } else {
                res = await axios.get(`${this.baseUrl}/${endpoint}`);
            }
            if (this.debug) {
                console.log(res)
            }
            //TODO: fix the api
            if (res.status === (200 | 202) ) {
                return res.data;
            }
        } catch (e) {
            console.warn(e);
        }
    }
}
