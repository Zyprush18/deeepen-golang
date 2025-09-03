import { redirect, type Actions } from "@sveltejs/kit";
import axios from "axios";
import * as z from "zod"

// schema register / validation for form registerd
const registerSchema =  z.object({
    username: z.string().min(3, "Username Must be 3 character") ,
    email: z.email("Not Email"),
    password: z.string().min(6, "Password Must be 6 character").max(12, "password max 12 character"),
    repassword: z.string().min(6, "Password Must be 6 character").max(12, "password max 12 character"),
}).refine((data) => data.password === data.repassword, {
    message: "Password don't match",
    path: ["repassword"]
})

export const actions: Actions = {
    default: async (event) => {
        const formDataa = Object.fromEntries(await event.request.formData());
        const talentData = registerSchema.safeParse(formDataa);

        console.log(talentData);
        // added to databases
        if (talentData.data != null) {
            await axios.post("http://localhost:3000/api/register",talentData.data).then((Response) =>{
                console.log(Response);
            }).catch((error)=>{
                console.log(error);
                
            })
            redirect(302,"/login")
        }

        
    }
};
