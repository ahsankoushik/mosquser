<script lang="ts">
    import { AtSign, Database, LogOut } from "@lucide/svelte";
    import Button from "../components/ui/button/button.svelte";
    import type MosqClient from "../api/client.svelte";
    import { Page } from "../api/client.svelte";
    import Users from "./Users.svelte";
    const {mc}:{mc:MosqClient} = $props();
    const page = mc.page;
    $effect(()=>{
      const url = new URL(window.location.href);
      const page = url.searchParams.get("page");
      switch (page){
        case "messenger":
          mc.page.set(Page.Messenger)
          break
        case "acls":
          mc.page.set(Page.Acls);
          break
        default:
          mc.page.set(Page.Users);
          break
      }
    })
</script>

<div class="w-screen h-screen min-h-[400px] flex">
    <div class="w-[5vw] min-w-[60px] flex-shrink-0">
        <div class="h-full flex flex-col justify-between">
            <div class=" grid py-4 gap-2">
                <div class="flex justify-center">
                    <Button
                        onclick={()=>{
                            mc.changePage(Page.Users)
                        }}
                        class="bg-white p-2 hover:bg-gray-200 {[Page.Users,Page.Acls].includes($page) ? "ring-2 ring-black" : ""}"
                    >
                        <Database color={"black"}/>
                    </Button>
                </div>
                <div class="flex justify-center">
                    <Button
                        onclick={()=>{
                            mc.changePage(Page.Messenger)
                        }}
                        class="bg-white p-2 hover:bg-gray-200 {$page == Page.Messenger ? "ring-2 ring-black" : ""}"
                    >
                        <AtSign color={"black"}/>
                    </Button>
                </div>
            </div>
    
            <div
                class="mb-6 flex justify-center"
            >
                <Button
                    onclick={()=>{
                        mc.logOut()
                    }}
                >
                    <LogOut />
                </Button>
            </div>
        </div>
    </div>
    <div class="w-[15vw] min-w-[200px] border flex-shrink-0">
        <div class="grid grid-cols-1">
            {#if $page == Page.Messenger}
                messenger Page
            {:else}
                <div class="grid grid-cols-1 p-3 gap-y-3">
                    <div class="">
                        <h2>Collections</h2>
                    </div>
                    <Button 
                        onclick={()=>{
                            mc.changePage(Page.Users)
                        }}
                        class="bg-white text-black text-center p-2 rounded-md hover:bg-gray-200 {$page == Page.Users ? "ring-2 ring-black" : ""}"
                    >
                        <p class="text-lg font-semibold">Users</p>
                    </Button>
                    <Button 
                        onclick={()=>{
                            mc.changePage(Page.Acls)
                        }}
                        class="bg-white text-black text-center p-2 rounded-md hover:bg-gray-200 {$page == Page.Acls ? "ring-2 ring-black" : ""}"
                    >
                        <p class="text-lg font-semibold">Acls</p>
                    </Button>
                </div>
            {/if}
        </div>
    </div>
    <div class="w-[80vw] min-w-[800px] ">
        {#if $page == Page.Acls}
            acl Page
        {:else if $page == Page.Messenger}
            Messenger
        {:else}
            <Users {mc}/>
        {/if}
    </div>
</div>