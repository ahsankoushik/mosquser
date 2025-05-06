interface MCRes {
    status: number,
    message: string,
    data: any
}

class MosqClient{
    host: string
    headers:HeadersInit
    loggedin = $state(false)
    loading = $state(false)
    private _token?:string


    constructor(){
        this.host = import.meta.env.VITE_HOST;
        this.headers = {
            "Content-Type" : "application/json", 
        };
    }
    set token(value:string){
        this._token = value;
    }
    get token():string|undefined{
        return this._token
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
            // mode: "cors",  // remove this when running with fiber
            ...(body && { body: JSON.stringify(body) })
        })
        const resJson = await res.json() as MCRes;
        return resJson
    }

    setAuthHeader(token:string){
        this.token = token;
        this.loggedin = true;
        localStorage.setItem("auth_token",token);
        this.headers = {
            ...this.headers,
            "Authorization": "Bearer " + token
        }
    }

    async login(email:string, password:string){
        this.loading = true
        const log =  await this.hit({
            path:"/auth/login",
            init:{
                method:"POST",
            },
            body:{
                email,
                password
            }
        })
        if(log.status==200){
            this.setAuthHeader(log.data.token)
        }
        this.loading = false;
        return log;
    }

    async checkAndRefreash(){
        const res = await this.hit({
            path:"/auth/refreash",
        });
        if (res.status == 200){
           this.setAuthHeader(res.data.token)
        }
    }
    async logOut(){
        this._token = undefined;
        this.loggedin = false;
        this.headers = {
            "Content-Type" : "application/json", 
        }
    }
}


export default MosqClient;