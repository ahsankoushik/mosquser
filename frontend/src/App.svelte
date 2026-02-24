<script>
    import { onMount } from "svelte";
  import MosqClient, { Page } from "./lib/api/client.svelte";
    import Loading from "./lib/comps/Loading.svelte";
  import Login from "./lib/pages/Login.svelte";
    import Main from "./lib/pages/Main.svelte";
    import  {Toaster} from 'svelte-5-french-toast'
  const mc = MosqClient.getInstance();
  onMount(()=>{
    const t = localStorage.getItem("auth_token");
    if(t != null){
      mc.token = t;
      mc.setAuthHeader(t);
      mc.checkAndRefreash();
    }
    const url = new URL(window.location.href);
    const page = url.searchParams.get("page");
    switch (page){
      case "messenger":
        mc.page.set(Page.Messenger)
        break
      case "acls":
        mc.page.set(Page.Acls);
        break
      case "broker":
        mc.page.set(Page.Broker);
        break
      default:
        mc.page.set(Page.Users);
        break
    }
  })
</script>

{#if mc.loading}
  <Loading text={""}/>
{/if}
{#if mc.loggedin}
  <Main {mc}/>
{:else}
  <Login {mc}/>
{/if}

<Toaster/>
