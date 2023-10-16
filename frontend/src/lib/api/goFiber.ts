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
        let res: any;
        try {
            if (id) {
                res = await axios.get(`${this.baseUrl}/${endpoint}/${id}`);
            } else {
                res = await axios.get(`${this.baseUrl}/${endpoint}`);
            }
            if (this.debug) {
                console.log(res)
            }
        } catch (e) {
            console.warn(e);
        } finally {
            return res.data
        }
    }

    async post(endpoint: string, payload: any, id?: string | number): Promise<any> {
        let res: any;
        try {
            if (id) {
                res = await axios.post(`${this.baseUrl}/${endpoint}/${id}`, payload);
            } else {
                res = await axios.post(`${this.baseUrl}/${endpoint}`, payload);
            }
        } catch (e) {
            console.warn(e);
        } finally {
            return res.data;
        }
    }
}
