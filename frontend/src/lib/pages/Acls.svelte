<script lang="ts">
    import type MosqClient from "../api/client.svelte";
    import Pagination from "../comps/Pagination.svelte";
    import type { MCResCol } from "../api/client.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { Plus } from "@lucide/svelte";
    import AclCreateModal from "../comps/AclCreateModal.svelte";

    const {mc}:{mc:MosqClient} = $props();
    let loading = $state(true)
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
        loading = false
    })

    let showDrawer = $state(false);

    const toggleDrawer = () => {
        showDrawer = !showDrawer;
    };
</script>

<div class="px-3 py-2 h-full flex flex-col justify-between">
    <div class="my-2 flex justify-between">
        <div>
            <h1 class="text-2xl font-bold">Acls</h1>
        </div>
        <div>
            <Button onclick={toggleDrawer}>
                <Plus />New Acl
            </Button>
        </div>
    </div>
    <div class="w-full grow overflow-scroll">
        <table class="table-auto w-full border-collapse border border-gray-300">
            <thead class="bg-gray-100 sticky top-0 shadow-md">
                <tr>
                    <th class="px-4 py-2 border border-gray-300 text-left">Email</th>
                    <th class="px-4 py-2 border border-gray-300 text-left">Super Users</th>
                    <th class="px-4 py-2 border border-gray-300 text-left">Created</th>
                    <th class="px-4 py-2 border border-gray-300 text-left">Updated</th>
                </tr>
            </thead>
            <tbody class="">
                {#if !loading}
                    {#each data.data as u, i (u.id)}
                        <tr class="hover:bg-gray-50 odd:bg-white even:bg-gray-50">
                            <td class="px-4 py-2 border border-gray-300">{u.email}</td>
                            <td class="px-4 py-2 border border-gray-300">{u.super_user}</td>
                            <td class="px-4 py-2 border border-gray-300">{new Date(u.created).toLocaleString()}</td>
                            <td class="px-4 py-2 border border-gray-300">{new Date(u.updated).toLocaleString()}</td>
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

{#if showDrawer}
    <AclCreateModal {toggleDrawer} {mc} {get}/>
{/if}

