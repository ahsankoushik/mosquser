<script>
  import MosqClient from "./lib/api/client.svelte";
    import Loading from "./lib/comps/Loading.svelte";
  import Login from "./lib/pages/Login.svelte";
    import Main from "./lib/pages/Main.svelte";
  const mc = new MosqClient();
  $effect(()=>{
    const t = localStorage.getItem("auth_token");
    if(t != null){
      mc.token = t;
      mc.setAuthHeader(t);
      mc.checkAndRefreash();
    }
  })

</script>

{#if mc.loading}
  <Loading/>
{/if}
{#if mc.loggedin}
  <Main/>
{:else}
  <Login {mc}/>
{/if}


