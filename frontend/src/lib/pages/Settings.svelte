<script>
    import MosqClient from "../api/client.svelte";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import SettingsSidebar from "../comps/SettingsSidebar.svelte";
    const mc = MosqClient.getInstance();
    
</script>



<div class="flex w-full h-full">
    <SettingsSidebar/>
    <div
        class="w-[80vw] min-w-[400px] border flex-shrink-0 p-2"
    >
        <h1 class="text-2xl font-semibold">Broker</h1>
        <div class="w-3/4 ">
            <div>
                <Label>Broker Url</Label>
                <Input/>
            </div>
        </div>
    </div>
</div>