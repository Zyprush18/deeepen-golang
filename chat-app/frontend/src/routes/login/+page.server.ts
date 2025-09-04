import z from 'zod';
import type { Actions } from './$types';
import axios from 'axios';
import { API_URL } from '$env/static/private';
import { redirect } from '@sveltejs/kit';

type Responapi = {
	status: number;
	Message: string;
	token: string;
};

const loginSchema = z.object({
	email: z.email('Field Email must be email'),
	password: z
		.string({
			error: (iss) => (iss.input === undefined ? 'Field is required.' : 'Invalid input.')
		})
		.min(6, 'Password Must be 6 character')
		.max(12, 'password max 12 character')
});

export const actions: Actions = {
	default: async (event) => {
		const formData = Object.fromEntries(await event.request.formData());
		const parsed = loginSchema.safeParse(formData);
		if (!parsed.success) {
			return parsed.error.message;
		}
        let response;

		try {
			response = await axios.post(API_URL + '/api/login', parsed.data);

			const apaaja: Responapi = {
				status: response.status,
				Message: response.data.message,
				token: response.data.token
			};
			if (apaaja.status === 200) {
				event.cookies.set('auth', apaaja.token, {
					path: '/chat',
					maxAge: 60 * 60 * 24 * 1,
					httpOnly: false
				});
			}
		} catch (error: any) {
            console.log(error);
            return {
                status: error.response.status,
                message: error.response.data.message
            }
		}
        
        if (response?.status === 200) {
            throw redirect(302,'/chat');
            
        }
	}
    
};
