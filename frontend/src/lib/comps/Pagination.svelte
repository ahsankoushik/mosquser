<script lang="ts">
    import { onMount } from "svelte";
    import type MosqClient from "../api/client.svelte";

    const {mc,totalPages,get}:{mc:MosqClient,totalPages:number,get:()=>Promise<void>} = $props();
    const pagination = mc.pagination;

    let pageCount = 5; // pages to view

    let firstPage = $state(false);
    let lastPage = $state(false);

    const changePage = (page:number, perPage:number)=>{
        pagination.set({
            page,
            perPage
        })
        get();
    }

    const pageListGen = ()=>{
		if(totalPages<pageCount){
			let arr:number[] = []
            for(let i = 1;i<=totalPages;i++){
                arr.push(i);
            }
            return arr;
		}
		let start = 1;
		let stop = totalPages;
        let leqr = Math.floor(pageCount/2);
		if($pagination.page<(Math.floor(pageCount)/2)+1){
            // less then first half
			stop = pageCount;
		}else if(!($pagination.page<(totalPages-Math.floor(pageCount/2)))){
            // less then second half
			start = totalPages - pageCount
		}else{
			start = $pagination.page - leqr;
			stop = $pagination.page + leqr;
		}
		let arr:number[] = []
		for(let i = start;i<=stop;i++){
			arr.push(i);
		}
        console.log(arr)
		return arr;
	}
	let listPages = $derived(pageListGen());

    $effect(()=>{
        if(!listPages.includes(1)){
        firstPage = true
        }
        if(!listPages.includes(totalPages)){
            lastPage = true
        }
    })
</script>

<div class="flex items-center justify-between mt-6">
    <button
        class="flex items-center px-5 py-2 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md gap-x-2 hover:bg-gray-100"
        disabled={$pagination.page==1}
        onclick={()=>{changePage($pagination.page-1,$pagination.perPage)}}
    >
        <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5 rtl:-scale-x-100"
        >
            <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18"
            />
        </svg>

        <span> previous </span>
    </button>

    <div class="items-center  lg:flex gap-x-3">
        {#if firstPage}
            <button  
                class="px-2 py-1 text-sm rounded-md hover:bg-gray-100 text-gray-500 "
                onclick={()=>{changePage(1,$pagination.perPage)}}
                >1
            </button>
            <p class="text-gray-500 text-sm">. . .</p>
        {/if}
        {#each listPages as page }
            <button  
                class="px-2 py-1 text-sm rounded-md {$pagination.page===page ? "bg-blue-100/60 text-blue-500" : "hover:bg-gray-100 text-gray-500"} "
                onclick={()=>{changePage(page,$pagination.perPage)}}
                >{page}
            </button>
        {/each}
        {#if lastPage}
            <p class="text-gray-500 text-sm">. . .</p>
            <button  
                class="px-2 py-1 text-sm rounded-md hover:bg-gray-100 text-gray-500 "
                onclick={()=>{changePage(totalPages,$pagination.perPage)}}
                >{totalPages}
            </button>
        {/if}
    </div>

    <button
        class=" flex items-center px-5 py-2 text-sm text-gray-700 capitalize transition-colors duration-200 bg-white border rounded-md gap-x-2 hover:bg-gray-100 disabled"
        disabled={$pagination.page==totalPages}
        onclick={()=>{changePage($pagination.page+1,$pagination.perPage)}}

    >
        <span> Next </span>

        <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5 rtl:-scale-x-100"
        >
            <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3"
            />
        </svg>
    </button>
</div>