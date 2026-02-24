<script>
    import MosqClient, { Page } from "../api/client.svelte";
    import Button from "../components/ui/button/button.svelte";
    const mc = MosqClient.getInstance();
    const page = mc.page;

</script>

<div class="w-[15vw] min-w-[200px] border flex-shrink-0">
    <div >
        <div class="grid grid-cols-1 p-3 gap-y-3">
            <div class="">
                <h2>Collections</h2>
            </div>
            <Button 
                onclick={()=>{
                    mc.changePage(Page.Broker)
                }}
                class="bg-white text-black text-center p-2 rounded-md hover:bg-gray-200 {$page == Page.Broker ? "ring-2 ring-black" : ""}"
            >
                <p class="text-lg font-semibold">Broker</p>
            </Button>
            <!-- <Button 
                onclick={()=>{
                    mc.changePage(Page.Acls)
                }}
                class="bg-white text-black text-center p-2 rounded-md hover:bg-gray-200 {$page == Page.Acls ? "ring-2 ring-black" : ""}"
            >
                <p class="text-lg font-semibold">Acls</p>
            </Button> -->
        </div>
    </div>
</div>