import { writable, type Readable, type Unsubscriber } from "svelte/store"
import { get } from 'svelte/store';

export interface MCRes {
    status: number,
    message: string,
    data: any
}
export interface MCResCol{
    status:number,
    message:string,
    page:number,
    perPage:number,
    totalPages:number,
    data:any
}
export enum Page{
    Users ="users",
    Acls = "acls",
    Messenger = "messenger",
    Login = "login",
    Broker = "broker"
    
}

export interface TokenData {
    exp : number,
    iat : number,
    id : number,
    username : string
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
    pagination = writable({
        page:1,
        perPage:30,
    })
    private _token?:string

    private static instance: MosqClient;

    static getInstance(): MosqClient {
        if (!MosqClient.instance) {
            MosqClient.instance = new MosqClient(); // Create if not already
        }
        return MosqClient.instance;
    }


    private constructor(){
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
    get tokenData():TokenData{

        return JSON.parse(atob(this._token!.split('.')[1])) as TokenData;
    }
    private async hit(
        {path,init,params,body}:{
            path:string,
            init?:RequestInit,
            params?:{[key:string]:any},
            body?:{[key:string]:any}
        }) : Promise<any>
    {
        let url = this.host + path;
        if(params != null){
            url += "?" + new URLSearchParams(params).toString();
        }
        try{
            const res =  await fetch(url,{
                ...init,
                headers:this.headers,
                // mode: "cors",  // remove this when running with fiber
                ...(body && { body: JSON.stringify(body) })
            })
            this.loading = false;
            const resJson = await res.json();
            return resJson
        }catch(e){
            console.log(e)
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
        }) as MCRes;
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
    async getItems(){
        const page = get(this.page)
        const pagination = get(this.pagination)
        this.addQueryParam("page",page);
        this.addQueryParam("pageNumber",pagination.page.toString());
        this.addQueryParam("perPage",pagination.perPage.toString());
        const res = await this.hit({
            path:`/${page}`,
            params:pagination
        }) as MCResCol;
        return res;
    }
    async createUser(email:string,password:string, super_user:boolean){
        const res = await this.hit({
            path:"/users",
            init:{
                method: "POST",
            },
            body : {
                email, 
                password,
                super_user
            }
        }) as MCRes;
        return res;
    }
    async updateUser(email:string,password:string="-", super_user:boolean){
        const res = await this.hit({
            path:"/users",
            init:{
                method: "PUT",
            },
            body : {
                email, 
                password,
                super_user
            }
        }) as MCRes;
        return res;
    }
    async deleteUser(id:number){
        const res = await this.hit({
            path:"/users",
            init:{
                method:"DELETE"
            },
            body:{
                id
            }
        }) as MCRes;
        return res
    }
    async userSearch(email:string){
        const res = await this.hit({
            path:"/users/search",
            params:{
                limit:5,
                email
            }
        }) as MCRes;
        return res;
    }
    async createAcl(userID:number,topic:string,acc:number){
        const res = await this.hit({
            path:"/acls",
            init:{
                method: "POST",
            },
            body : {
                userID,
                topic,
                acc
            }
        }) as MCRes;
        return res;
    }
    async updateAcl(ID:number,topic:string,acc:number){
        const res = await this.hit({
            path:"/acls",
            init:{
                method: "PUT",
            },
            body : {
                ID,
                topic,
                acc
            }
        }) as MCRes;
        return res;
    }
    async deleteAcl(id:number):Promise<MCRes> {
       console.log(id);
       const res = await this.hit({
            path: "/acls",
            init: {
                method: "DELETE"
            },
            body: {id}
       }) as MCRes;
       return res;
    }
    
    async getKeys(keys:Array<string>):Promise<MCRes>{
        const rest = await this.hit({
            path : "/key_value",
            params : {
                keys:keys
            }
        }) as MCRes
        return rest
    }
    async setKey(key:string, value:string):Promise<MCRes>{
        const res = await this.hit({
            path:"/key_value",
            init: {
                method: "POST",
            },
            body : {
                key,
                value
            }
        }) as MCRes;
        return res;
    }
}

export default MosqClient;
