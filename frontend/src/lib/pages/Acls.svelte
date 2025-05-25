<script lang="ts">
    import type MosqClient from "../api/client.svelte";
    import Pagination from "../comps/Pagination.svelte";
    import { Page, type MCResCol } from "../api/client.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { Plus } from "@lucide/svelte";
    import AclCreateModal from "../comps/AclCreateModal.svelte";
    import DbSidebar from "../comps/DBSidebar.svelte";

    const {mc}:{mc:MosqClient} = $props();
    const page = mc.page;
    let loading = $state(true);
    let showDrawer = $state(false);
    let update = $state(-1);
    let data:MCResCol = $state({
        status:-1,
        message:"",
        page:1,
        perPage:30,
        totalPages:1,
        data:{}
    })
    const get = async()=>{
        data = await mc.getItems();
    }
    $effect(()=>{
        const url = new URL(window.location.href);
        let pageNumber = parseInt(url.searchParams.get("pageNumber") || "1");
        if(isNaN(pageNumber)) {pageNumber = 0};
        let perPage = parseInt(url.searchParams.get("perPage") || "30");
        if(isNaN(perPage)){perPage = 30}
        mc.pagination.set({
            page: pageNumber,
            perPage
        })
        get();
        loading = false
    })


    const toggleDrawer = () => { 
        update = -1;
        showDrawer = !showDrawer;
    };
    $inspect(data)
</script>
<svelte:head>
    <title>Acls | Mosquser</title>
</svelte:head>
<div class="flex w-full h-full">
    <DbSidebar {mc} />
    <div class="w-[80vw] min-w-[400px] border flex-shrink-0">
        <div class=" h-full flex flex-col justify-between">
            <div class="m-2 flex justify-between">
                <div>
                    <h1 class="text-2xl font-bold">Acls</h1>
                </div>
                <div>
                    <Button onclick={toggleDrawer}>
                        <Plus />New Acl
                    </Button>
                </div>
            </div>
            <div class="w-full grow overflow-y-scroll p-2">
                <table class="table-auto w-full border-collapse border border-gray-300">
                    <thead class="bg-gray-100 sticky top-0 shadow-md">
                        <tr>
                            <th class="px-4 py-2 border border-gray-300 text-left">User</th>
                            <th class="px-4 py-2 border border-gray-300 text-left">Topic</th>
                            <th class="px-4 py-2 border border-gray-300 text-left">Acc</th>
                            <th class="px-4 py-2 border border-gray-300 text-left">Created</th>
                            <th class="px-4 py-2 border border-gray-300 text-left">Updated</th>
                        </tr>
                    </thead>
                    <tbody class="">
                        {#if !loading}
                            {#each data.data as device, i (device.id)}
                                <tr class="hover:bg-gray-50 odd:bg-white even:bg-gray-50"
                                    onclick={()=>{
                                        update = i;
                                        showDrawer = true;
                                    }}
                                >
                                    <td class="px-4 py-2 border border-gray-300">{device.user.email}</td>
                                    <td class="px-4 py-2 border border-gray-300">{device.topic}</td>
                                    <td class="px-4 py-2 border border-gray-300">{device.acc == 3 ? "Read & Write" : device.acc == 2 ? "Write" : "Read"}</td>
                                    <td class="px-4 py-2 border border-gray-300">{new Date(device.created).toLocaleString()}</td>
                                    <td class="px-4 py-2 border border-gray-300">{new Date(device.updated).toLocaleString()}</td>
                                </tr>
                            {/each}
                        {/if}
                    </tbody>
                </table>
            </div>
            <div class="z-1 ">
                <Pagination {mc} totalPages={data.totalPages} get={get} />
            </div>
        </div>
    </div>
</div>


{#if showDrawer}
    <AclCreateModal {toggleDrawer} {mc} {get} {update} {data}/>
{/if}

