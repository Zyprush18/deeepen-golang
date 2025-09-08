import { getData } from '$lib/utils/profile.js';
import { redirect } from '@sveltejs/kit';


export async function load({ cookies,params }) {
    const slug = params.slug    
    const auth = cookies.get('auth');

    if (auth === '' || auth === undefined) {
        throw redirect(302, '/login');
    }

    const profile = await getData(auth);
    
    if (profile.status != 200) {	
        cookies.delete("auth", {path:"/chat"});
        throw redirect(302, "/login");
        
    }
    return {Auth: auth, Profile: profile,Slug: slug}
}
