<script lang="ts">
    import { slide } from "svelte/transition";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Checkbox from "../components/ui/checkbox/checkbox.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { Save, UserSearch } from "@lucide/svelte";
    import type MosqClient from "../api/client.svelte";
    import type { MCRes } from "../api/client.svelte";
    const { toggleDrawer,mc, get}:{toggleDrawer:()=>void,mc:MosqClient,get:()=>Promise<void>} =  $props();
    let email = $state("")
    let topic = $state("")
    let acc = $state(0)
    let openSerachHelper = $state(false)
    let userID = -1;
    let timeOutId:number;
    let searchData:MCRes = $state({
        status:-1,
        message:"",
        data:[]
    })
</script>

<div
    class="fixed top-0 right-0 h-full w-1/2 min-w-[300px] bg-white shadow-lg transition-transform transform translate-x-0"
    style="z-index: 1000;"
    transition:slide={{
        axis:"x",
        duration:100
    }}
>
    <div class="p-4 flex justify-between items-center border-b border-gray-300">
        <h2 class="text-xl font-bold">New User</h2>
        <button class="text-gray-500 hover:text-gray-700" onclick={toggleDrawer}>
            ✕
        </button>
    </div>
    <form 
        action=""
        autocomplete="off"
        onsubmit={async(e)=>{
            e.preventDefault();
            
        }}
    >
        <div class="p-4 grid grid-cols-1 gap-3">
            <!-- Add your form or content for the drawer here -->
            <div>
                <Label for="email">Email:</Label>
                <Input
                    bind:value={email}
                    id="email"
                    type="email"
                    placeholder="Enter Email address here"
                    required
                    oninput={(e)=>{
                        if(timeOutId != undefined){
                            clearTimeout(timeOutId);
                        }
                        userID = -1;
                        timeOutId = setTimeout(async()=>{
                            searchData = await mc.userSearch(email);
                        },1000)
                    }}
                    on:focusin={()=>{
                        openSerachHelper = true;
                    }}
                    on:focusout={()=>{
                        setTimeout(()=>{
                            openSerachHelper = false;
                        },200)
                    }}
                />
                {#if openSerachHelper}
                    <ul class=" mx-2 mt-2">
                        {#each searchData.data as user (user.id) }
                            <!-- svelte-ignore a11y_click_events_have_key_events -->
                            <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
                            <li 
                                class="px-3 hover:bg-blue-600 hover:text-white cursor-pointer"
                                onclick={()=>{
                                    userID = user.id;
                                    email = user.email;
                                }}
                            >{user.email}</li>
                        {/each}
                    </ul>
                {/if}
            </div>
            <div>
                <Label for="topic">Topic:</Label>
                <Input
                    bind:value={topic}
                    id="topic"
                    type="text"
                    placeholder="Enter topic name here"  
                />
            </div>
            <div
            >
                <Label for="acc">Acc:</Label>
                <Input
                    bind:value={topic}
                    id="acc"
                    type="number"
                    placeholder="Enter Acc value here"  
                />
            </div>
            <Button
                type="submit"
            >
                <Save class="mr-2"/>
                Save
            </Button>
        </div>
    </form>
</div>
<div
    class="fixed inset-0 bg-black bg-opacity-50"
    style="z-index: 999;"
    onclick={toggleDrawer}
></div>