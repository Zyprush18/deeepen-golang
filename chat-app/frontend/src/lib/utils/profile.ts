import axios from 'axios';
import { API_URL } from '$env/static/private';

type friend = {
	from_id: string,
    to_user_id: string,
    name_from: string,
    name_to: string,
    uuid: string,
    status: string,
    role: string
};

type Response = {
    status: number
    username:string
    email:string
    uuid:string
}

export async function getDataProfile(auth: string): Promise<Response> {
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


export async function getDataFriends(auth: string): Promise<friend> {
	console.log(auth);
	
	try {
		const response  = await axios.get(`${API_URL}/api/friend`, {
			headers:{
				"Content-Type":"application/json",
				Authorization: "Bearer "+auth
			}
		})


		return response.data as friend
	} catch (error: any) {
		console.log(error);
		throw error.response
	}
}