import { getData } from '$lib/utils/profile.js';
import { redirect } from '@sveltejs/kit';


export async function load({ cookies }) {
	const auth = cookies.get('auth');

	if (auth === '' || auth === undefined) {
		throw redirect(302, '/login');
	}

	const profile = await getData(auth)
    return {Auth: auth, Profile: profile}
}
