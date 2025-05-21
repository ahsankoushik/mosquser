<script lang="ts">
    import { AtSign, Database, LogOut, Settings as SettingsIcon } from "@lucide/svelte";
    import Button from "../components/ui/button/button.svelte";
    import type MosqClient from "../api/client.svelte";
    import { Page } from "../api/client.svelte";
    import Users from "./Users.svelte";
    import Acls from "./Acls.svelte";
    import Messenger from "./Messenger.svelte";
    import Settings from "./Settings.svelte";
    const {mc}:{mc:MosqClient} = $props();
    const page = mc.page;
    
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
                <div class="flex justify-center">
                    <Button
                        onclick={()=>{
                            mc.changePage(Page.Broker)
                        }}
                        class="bg-white p-2 hover:bg-gray-200 {$page == Page.Broker ? "ring-2 ring-black" : ""}"
                    >
                        <SettingsIcon color={"black"}/>
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
    
    <div class="w-[95vw] min-w-[800px] ">
        {#if $page == Page.Acls}
            <Acls {mc}/>
        {:else if $page == Page.Messenger}
            <Messenger {mc}/>
        {:else if $page == Page.Broker}
            <Settings/>
        {:else}
            <Users {mc}/>
        {/if}
    </div>
</div>