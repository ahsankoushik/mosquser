<script lang="ts">
    import { slide } from "svelte/transition";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Checkbox from "../components/ui/checkbox/checkbox.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { Save } from "@lucide/svelte";
    import type MosqClient from "../api/client.svelte";
    import type { HTMLFormAttributes } from "svelte/elements";
    const { toggleDrawer,mc, get}:{toggleDrawer:()=>void,mc:MosqClient,get:()=>Promise<void>} =  $props();
    let email = $state("")
    let password = $state("")
    let superUser = $state(false)
    let formElement: HTMLFormElement | undefined;
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
            const res = await mc.createUser(email,password,superUser);
            console.log(res);
            if(res.status == 200){
                if(formElement != undefined){
                    formElement.reset();
                }
                get();
            }
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
                    autocomplete="none"
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