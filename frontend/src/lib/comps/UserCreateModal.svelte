<script lang="ts">
    import { slide } from "svelte/transition";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Checkbox from "../components/ui/checkbox/checkbox.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { Save, Trash2 } from "@lucide/svelte";
    import type MosqClient from "../api/client.svelte";
    import type { HTMLFormAttributes } from "svelte/elements";
    import type { MCResCol } from "../api/client.svelte";
    import toast from "svelte-5-french-toast";
    const { 
        toggleDrawer, 
        mc, 
        get,
        update,
        data,
    }:{
        toggleDrawer:()=>void,
        mc:MosqClient,
        get:()=>Promise<void>,
        update:number,
        data:MCResCol
    } =  $props();
    let id = $state(-1)
    let username = $state("")
    let password = $state("")
    let superUser = $state(false)
    let formElement: HTMLFormElement | undefined;
    $effect(()=>{
        if(update>-1){
            id = data.data[update].id;
            username = data.data[update].email;
            password = data.data[update].password;
            superUser = data.data[update].super_user
        }else{
            id = -1;
            username = "";
            password = "";
            superUser = false;
        }
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
        bind:this= {formElement}
        onsubmit={async(e)=>{
            e.preventDefault();
            if(update == -1){
                const res = await mc.createUser(username,password,superUser);
                if(res.status == 200){
                    if(formElement != undefined){
                        formElement.reset();
                    }
                    get();
                    toast.success("Added");
                }else{
                    toast.error(res.message)
                }
            }else{
                console.log("its working")
                const res = await mc.updateUser(id,username,password,superUser);
                if(res.status == 200){
                    if(formElement != undefined){
                        formElement.reset();
                    }
                    get();
                    toast.success("Updated")
                }else{
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
                    placeholder="Enter username here"
                    required
                    autocomplete="off"
                />
            </div>
            <div>
                <Label for="password">Password:</Label>
                <Input
                    bind:value={password}
                    id="password"
                    type="password"
                    placeholder="*************"  
                    autocomplete="new-password"
                />
            </div>
            <div
                class="flex gap-x-2 items-center"
            >
                <Checkbox id="admin" bind:checked={superUser}/>
                <Label for="admin">Is Admin</Label>
            </div>
            <div class="flex gap-2">
                <Button
                    class="flex-grow"
                    type="submit"
                >
                    <Save class="mr-2"/>
                    Save
                </Button>
                {#if update > -1}
                    <Button
                        class="bg-red-600 flex-grow"
                        onclick={async()=>{
                            const res = await mc.deleteUser(data.data[update].id);
                            if(res.status == 200){
                                get()
                                toggleDrawer();
                            }else{
                                // TODO :: show modal
                            }
                        }}
                    >
                        <Trash2 class="mr-2"/>
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
