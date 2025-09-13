import { getDataFriends, getDataProfile } from '$lib/utils/profile.js';
import { redirect } from '@sveltejs/kit';


export async function load({ cookies,params,depends }) {
    const slug = params.slug;
    const auth = cookies.get('auth');

    if (auth === '' || auth === undefined) {
        throw redirect(302, '/login');
    }

    const profile = await getDataProfile(auth);
    const friends = await getDataFriends(auth);
    
    if (profile.status != 200) {	
        cookies.delete("auth", {path:"/chat"});
        throw redirect(302, "/login");
        
    }
    depends(slug);



    return {Auth: auth, Profile: profile,Slug: slug, Friend: friends }
}
