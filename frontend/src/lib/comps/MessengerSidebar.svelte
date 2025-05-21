<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import type MosqClient from "../api/client.svelte";
    import { Topics } from "../api/stores";
    import { get } from "svelte/store";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { Save } from "@lucide/svelte";

    
    const {mc,subscribe} : {mc:MosqClient,subscribe:(topic:string)=>void} = $props();
    let topic = $state("");


    // onMount(()=>{
    //     const topics = localStorage.getItem("mqtt_topics") 
    //     if(topics != null){
    //         Topics.set(JSON.parse(topics))
    //     }
    // })
    const handleSave = ()=>{
        if(topic == ""){
            return
        }
        subscribe(topic);
        Topics.update((t)=>{
            t.push({
                topic,
                connected:true
            });
            return t
        });
        // let topics = localStorage.getItem("mqtt_topics");
        // let topicsArr = [];
        // if(topics != null){
        //     topicsArr = JSON.parse(topics);
        // }
        // topicsArr.push(topic);
        
        topic = "";
    }   
</script>
<div class="w-[25vw] min-w-[200px] border flex-shrink-0">
    <div class=" h-min-[200px] px-2 py-1 gap-y-1 ">
        <div>
            <Label for="topic">Topic</Label>
            <div class="flex gap-1">
                <Input
                    bind:value={topic}
                    id="topic"
                    onkeydown={(e)=>{
                        if(e != null){
                            if(e.key == "Enter"){
                                handleSave();
                            }
                        }
                    }}
                />
                <Button
                    onclick={handleSave}
                >
                    <Save/>
                </Button>
            </div>
            
        </div>
        {#each $Topics as tp,i}
            <div
                class="my-1 text-center odd:bg-gray-200 odd:bg-opacity-60 p-2 rounded-lg hover:bg-gray-400 cursor-pointer"
            >
                {tp.topic}
            </div>
        {/each}
    </div>
</div>
