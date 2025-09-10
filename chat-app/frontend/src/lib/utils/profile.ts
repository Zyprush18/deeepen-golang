import axios from 'axios';
import { API_URL } from '$env/static/private';

type friend = {
	name: string;
	status: string;
	uuid: string;
};

type Response = {
    status: number
    username:string
    email:string
    uuid:string
	data: {
		friend: friend[]
	}
}

export async function getData(auth: string): Promise<Response> {
	try {
        const resp = await axios.get(`${API_URL}/api/profile`, {
			headers: {
				Authorization: 'Bearer ' + auth
			}
		});

        return resp.data as Response;
	} catch (error: any) {
		console.log(error);
        throw error.response
	}
}
