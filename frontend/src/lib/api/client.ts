interface MCRes {
    status: number,
    message: string,
    data: any
}

class MosqClient{
    host: string
    headers:HeadersInit

    constructor(){
        this.host = import.meta.env.VITE_HOST;
        this.headers = {
            "Content-Type" : "application/json", 
        };
    }
    
    
    private async hit(
        {path,init,params,body}:{
            path:string,
            init?:RequestInit,
            params?:URLSearchParams,
            body?:{[key:string]:any}
        }) : Promise<MCRes>
    {
        let url = this.host + path;
        if(params != null){
            url += "?" + params.toString();
        }
        const res =  await fetch(url,{
            ...init,
            headers:this.headers,
            body:JSON.stringify(body)
        })
        const resJson = await res.json() as MCRes;
        return resJson
    }

    async login(email:string, password:string){
        return await this.hit({
            path:"/auth/login",
            init:{
                method:"POST",
                credentials: "include"
            },
            body:{
                email,
                password
            }
        })
    }
    
}


export default MosqClient;