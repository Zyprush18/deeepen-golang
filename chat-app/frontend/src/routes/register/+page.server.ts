import { API_URL } from '$env/static/private';
import { redirect, type Actions } from '@sveltejs/kit';
import axios from 'axios';
import * as z from 'zod';

// schema register / validation for form registerd
const registerSchema = z
	.object({
		username: z
			.string({
				error: (iss) => (iss.input === undefined ? 'Field is required.' : 'Invalid input.')
			})
			.min(3, 'Username Must be 3 character'),
		email: z.email('Field Email must be email'),
		password: z
			.string({
				error: (iss) => (iss.input === undefined ? 'Field is required.' : 'Invalid input.')
			})
			.min(6, 'Password Must be 6 character')
			.max(12, 'password max 12 character'),
		repassword: z
			.string({
				error: (iss) => (iss.input === undefined ? 'Field is required.' : 'Invalid input.')
			})
			.min(6, 'Repassword Must be 6 character')
			.max(12, 'repassword max 12 character')
	})
	.refine((data) => data.password === data.repassword, {
		message: "password and repassword don't match",
		path: ['repassword']
	});

export const actions: Actions = {
	default: async (event) => {
		const formData = Object.fromEntries(await event.request.formData());
		const parsed = registerSchema.safeParse(formData);
		if (!parsed.success) {
			console.log(parsed.error.message);
			return parsed.error.issues;
		}

		try {
			const response = await axios.post(API_URL + '/api/register', parsed.data);
			console.log(response.data);
            return response.data
		} catch (error: any) {
			console.log(error);
            return {
                status: error.response?.data?.status,
                message: error.response?.data?.message
            }
		}

		// redirect(302,"/login")
	}
};
