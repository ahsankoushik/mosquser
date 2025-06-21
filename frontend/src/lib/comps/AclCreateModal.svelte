<script lang="ts">
    import { slide } from "svelte/transition";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { Save, Trash2 } from "@lucide/svelte";
    import  MosqClient from "../api/client.svelte";
    import type { MCRes } from "../api/client.svelte";
    import * as Select from "../components/ui/select";
    import toast from "svelte-5-french-toast";
    
    const { toggleDrawer,
            mc, 
            get,
            update, 
            data
        }:{
            toggleDrawer:()=>void,
            mc:MosqClient,
            get:()=>Promise<void>,
            update:number,
            data:MCRes} =  $props();
    const Accs = [
        {label:"Read",value:1},
        {label:"Write",value:2},
        {label:"Read & Write",value:3},
    ]
    let username = $state("")
    let topic = $state("")
    let acc = $state(Accs[0])
    let openSerachHelper = $state(false)
    let userID = $state(-1);
    let id = $state(-1);
    let timeOutId:NodeJS.Timeout;
    let searchData:MCRes = $state({
        status:-1,
        message:"",
        data:[]
    })

    $effect(()=>{
        if(update>-1){
            username = data.data[update].user.username;
            topic = data.data[update].topic;
            acc = Accs.filter((a)=>a.value == data.data[update].acc)[0];
            userID = data.data[update].user.id;
            id = data.data[update].id;
        }else{
            username = "";
            topic = "";
            acc = Accs[0];
            userID = -1;
            id = -1;
        }
    })
    const handleDelete = async()=>{
        console.log("ok")
        const res = await mc.deleteAcl(id);
        if(res.status == 200){
            get(); 
            toggleDrawer();
        }
    }
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
        <h2 class="text-xl font-bold">New Acl Record</h2>
        <button class="text-gray-500 hover:text-gray-700" onclick={toggleDrawer}>
            ✕
        </button>
    </div>
    <form 
        action=""
        autocomplete="off"
        onsubmit={async(e)=>{
            e.preventDefault();
            if(update == -1){
                const data = await mc.createAcl(userID,topic,acc.value);
                console.log(data);
                if(data.status == 200){
                    get();
                    toast.success("Added");
                }else if(data.status == 409){
                    toast.error(data.message);
                }
            }else{
                const data = await mc.updateAcl(id,topic,acc.value);
                if(data.status == 200){
                    toggleDrawer();
                    get();
                }else if(data.status == 409){
                    toast.error(data.message);
                }
            }
        }}
    >
        <div class="p-4 grid grid-cols-1 gap-3">
            <!-- Add your form or content for the drawer here -->
            <div>
                <Label for="username">Username:</Label>
                <Input
                    bind:value={username}
                    id="username"
                    type="username"
                    placeholder="Enter username address here"
                    required
                    oninput={(e)=>{
                        if(timeOutId != undefined){
                            clearTimeout(timeOutId);
                        }
                        userID = -1;
                        timeOutId = setTimeout(async()=>{
                            searchData = await mc.userSearch(username);
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
                                    username = user.username;
                                }}
                            >{user.username}</li>
                        {/each}
                    </ul>
                {/if}
                {#if userID == -1}
                    <p class="font-bold text-xs text-red-600">You Have to select a user.</p>
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
                <Select.Root portal={null} 
                    bind:selected={acc}
                >
                    <Select.Trigger class="w-full">
                      <Select.Value placeholder="Select Acc Value" />
                    </Select.Trigger>
                    <Select.Content>
                      <Select.Group>
                        {#each Accs as acc}
                          <Select.Item value={acc.value} label={acc.label}
                            >{acc.label}</Select.Item
                          >
                        {/each}
                      </Select.Group>
                    </Select.Content>
                    <Select.Input name="favoriteFruit" required />
                  </Select.Root>
            </div>
            <div
                class="flex gap-2"
            >
            <Button
                class="flex-1"
                type="submit"
            >
                <Save class="mr-2"/>
                Save
            </Button>
            {#if update != -1 }
                <Button
                    class="flex-1 bg-red-600"
                    onclick={handleDelete}
                >
                    <Trash2/>
                    Delete 
                </Button>
                
            {/if}
                
            </div>
        </div>
    </form>
</div>
<div
    class="fixed inset-0 bg-black bg-opacity-50"
    style="z-index: 999;"
    onclick={toggleDrawer}
></div>
