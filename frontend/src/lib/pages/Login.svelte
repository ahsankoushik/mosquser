<script lang="ts">
    import{ LogIn,  ShieldPlus } from "@lucide/svelte";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Button from "../components/ui/button/button.svelte";
    import type MosqClient from "../api/client.svelte";
    const {mc}:{mc:MosqClient} = $props();
    let username = $state("")
    let password = $state("")
    const error = mc.error;

</script>

<div class="h-screen flex justify-center items-center ">
    <div class="w-full sm:w-1/2 lg:w-1/3 mx-16 sm:mx-0">
        <div class=" flex justify-center text-[38px] font-bold gap-2 items-center">
            <ShieldPlus size={46} color={"purple"}/>
            <h1 class="text-center">MosqUser</h1>
        </div>
        <div class="my-2">
            <h2 class="text-xl text-center">
                Admin Login
            </h2>
        </div>
        <form action=""
            onsubmit={async(e)=>{
                e.preventDefault();
                const data = await mc.login(username,password);
            }}
        >
            <div class="my-2">
                <Label class="{$error.error ? "text-red-600" : ""}">Username</Label>
                <Input 
                    class="w-full {$error.error ? "ring ring-red-600" : ""}"
                    placeholder="Enter your username here"
                    required
                    bind:value={username}
                />
            </div>
            <div class="my-2">
                <Label class="{$error.error ? "text-red-600" : ""}">Password</Label>
                <Input 
                    class="w-full {$error.error ? "ring ring-red-600" : ""}"
                    placeholder="Enter password here"
                    type="password"
                    required
                    bind:value={password}
                />
            </div>
            {#if $error.error}
                <p class="text-sm text-red-600">{$error.message}</p>
            {/if}

            <Button
                class="mt-3 w-full gap-2 group"
                type="submit"
            >
                Login <LogIn
                    class="transition-transform duration-300 ease-in-out group-hover:translate-x-3"
                />
            </Button>
        </form>
    </div>
</div>
