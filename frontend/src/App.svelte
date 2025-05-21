<script>
    import { onMount } from "svelte";
  import MosqClient, { Page } from "./lib/api/client.svelte";
    import Loading from "./lib/comps/Loading.svelte";
  import Login from "./lib/pages/Login.svelte";
    import Main from "./lib/pages/Main.svelte";
  const mc = new MosqClient();
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
      case "settings":
        mc.page.set(Page.Settings);
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


