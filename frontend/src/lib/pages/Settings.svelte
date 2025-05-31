<script>
    import { onMount } from "svelte";
    import MosqClient from "../api/client.svelte";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import SettingsSidebar from "../comps/SettingsSidebar.svelte";
    import Button from "../components/ui/button/button.svelte";
    const mc = MosqClient.getInstance();
    let brokerUrl = $state("");
    onMount(async()=>{
        const res = await mc.getKeys(["broker_url"]);
        if(res.data.length > 0){
            for(let i = 0; i < res.data.length; i++){
                switch (res.data[i].key){
                    case "broker_url":
                        brokerUrl = res.data[i].value;
                }
            }
        }
    })

</script>

<svelte:head>
    <title>Settings | Mosquser</title>
</svelte:head>

<div class="flex w-full h-full">
    <SettingsSidebar/>
    <div
        class="w-[80vw] min-w-[400px] border flex-shrink-0 p-2"
    >
        <h1 class="text-2xl font-semibold">Broker</h1>
        <div class="w-3/4 ">
            <div>
                <Label>Broker Url</Label>
                <Input
                    bind:value={brokerUrl}
                />
            </div>
        </div>
        <div
            class="my-2"
        >
            <Button
                class=""
                onclick={async()=>{
                    mc.setKey("broker_url",brokerUrl);
                }}
            >
                Save
            </Button>
        </div>
    </div>
</div>
