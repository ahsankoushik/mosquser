import { writable, type Readable } from "svelte/store"

interface MCRes {
    status: number,
    message: string,
    data: any
}

export enum Page{
    Users ="users",
    Acls = "acls",
    Messenger = "messenger",
    Login = "login"
}



class MosqClient{
    host: string
    headers:HeadersInit
    loading = $state(false)
    loggedin = $state(false)
    error = writable({
        error:false,
        message:""
    })
    page = writable<Page>(Page.Users)
    private _token?:string


    constructor(){
        this.host = import.meta.env.VITE_HOST;
        this.headers = {
            "Content-Type" : "application/json", 
        };
    }
    set token(value:string){
        this.error.update((err)=>{
            err.error = false,
            err.message = ""
            return err
        })
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
        try{
            const res =  await fetch(url,{
                ...init,
                headers:this.headers,
                // mode: "cors",  // remove this when running with fiber
                ...(body && { body: JSON.stringify(body) })
            })
            this.loading = false;
            const resJson = await res.json() as MCRes;
            return resJson
        }catch(e){
            this.loading = false;
            const resJson:MCRes = {
                status:600,
                message:"unable to connect to the server.",
                data:{}
            }
            return resJson
        }


    }

    setAuthHeader(token:string){
        this.loggedin = true;
        this.token = token;
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
            this.changePage(Page.Users);
            this.setAuthHeader(log.data.token);
        }else{
            this.error.update((err)=>{
                err.error = true,
                err.message = log.message
                return err
            })
        }
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
        localStorage.removeItem("auth_token");
    }
    addQueryParam (key:string, value:string)  {
        const url = new URL(window.location.href);
        url.searchParams.set(key, value);
        window.history.pushState({}, '', url.toString());
    };
    changePage (page:Page){
        this.addQueryParam("page",page);
        this.page.set(page);
    }

}


export default MosqClient;