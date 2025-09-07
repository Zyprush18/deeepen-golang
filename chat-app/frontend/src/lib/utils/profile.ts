import axios from 'axios';
import { API_URL } from '$env/static/private';

type Response = {
    status: number
    username:string
    email:string
    uuid:string
}

export async function getData(auth: string): Promise<Response> {
	try {
        const resp = await axios.get(`${API_URL}/api/profile`, {
			headers: {
				Authorization: 'Bearer ' + auth
			}
		});

        return resp.data as Response
	} catch (error) {
		console.log(error);
        throw error
	}
}
