import { getDataFriends, getDataProfile } from '$lib/utils/profile.js';
import { redirect } from '@sveltejs/kit';


export async function load({ cookies }) {
	const auth = cookies.get('auth');

	if (auth === '' || auth === undefined) {
		throw redirect(302, '/login');
	}

	// ambil data profile
	const profile = await getDataProfile(auth);
	
	// ambil data friends
	const friends =  await getDataFriends(auth);

	if (profile.status != 200) {	
		cookies.delete("auth", {path:"/chat"});
		throw redirect(302, "/login");
		
	}
    return {Auth: auth, Profile: profile, Friend: friends}
}
