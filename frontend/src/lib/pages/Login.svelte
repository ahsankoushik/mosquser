<script>
    import { LogIn, ShieldX } from "@lucide/svelte";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Button from "../components/ui/button/button.svelte";
    import MosqClient from "../api/client";
    const host = import.meta.env.VITE_HOST
    let email = $state("")
    let password = $state("")

</script>

<div class="h-screen flex justify-center items-center ">
    <div class="w-full sm:w-1/2 lg:w-1/3 mx-16 sm:mx-0">
        <div class=" flex justify-center text-[38px] font-bold gap-2 items-center">
            <ShieldX size={46} color={"purple"}/>
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
                const mc = new MosqClient()
                const data = await mc.login(email,password);
                console.log(data)
            }}
        >
            <div class="my-2">
                <Label>Email</Label>
                <Input 
                    class="w-full"
                    placeholder="Enter your email here"
                    type="email"
                    required
                    bind:value={email}
                />
            </div>
            <div class="my-2">
                <Label>Password</Label>
                <Input 
                    class="w-full"
                    placeholder="Enter password here"
                    type="password"
                    required
                    bind:value={password}
                />
            </div>
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