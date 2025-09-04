import { redirect } from '@sveltejs/kit';

export function load({cookies}){
    const auth = cookies.get('auth');
    
    if (auth === '' || auth === undefined) {
        throw redirect(302, '/login');
    }
    
    return  {
        Auth: auth  
    }
}