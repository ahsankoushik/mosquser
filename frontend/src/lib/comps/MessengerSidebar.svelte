<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import type MosqClient from "../api/client.svelte";
    import { Topics } from "../api/stores";
    import { get } from "svelte/store";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { BadgePlus, Save, Unplug } from "@lucide/svelte";
    import ButtonWithHelp from "./ButtonWithHelp.svelte";

    const {
        mc,
        subscribe,
        unsubscribe,
    }: {
        mc: MosqClient;
        subscribe: (topic: string) => void;
        unsubscribe: (topic: string) => void;
    } = $props();
    let topic = $state("");

    // onMount(()=>{
    //     const topics = localStorage.getItem("mqtt_topics")
    //     if(topics != null){
    //         Topics.set(JSON.parse(topics))
    //     }
    // })
    const handleSave = () => {
        if (topic == "") {
            return;
        }
        subscribe(topic);
        Topics.update((t) => {
            t.push({
                topic,
                connected: true,
            });
            return t;
        });
        topic = "";
    };
    const handleUnsubscribe = (topic: string) => {
        if (topic == "") {
            return;
        }
        unsubscribe(topic);
        Topics.update((t) => {
            t = t.filter((e) => {
                return e.topic != topic;
            });
            return t;
        });
    };
    const handleSaveToLocal = ()=>{
        if(topic ==""){
            return
        }
        let topics = localStorage.getItem("mqtt_topics");
        let topicsArr = [];
        if(topics != null){
            topicsArr = JSON.parse(topics);
        }
        topicsArr.push(topic);
        localStorage.setItem("mqtt_topics", JSON.stringify(topicsArr));
    }
</script>

{#snippet subscribeElement()}
    <Button onclick={handleSave}>
        <BadgePlus />
    </Button>
{/snippet}
{#snippet saveToLocal({topic})}
    <button class="hover:bg-gray-300 p-1 rounded-lg" onclick={() => {}}
        ><Save /></button
    >
{/snippet}

{#snippet unsubscribeElement({topic})}
    <button
        class="hover:bg-gray-300 p-1 rounded-lg"
        onclick={() => {
            handleUnsubscribe(topic);
        }}><Unplug /></button
    >
{/snippet}

<div class="w-[25vw] min-w-[200px] border flex-shrink-0">
    <div class=" h-min-[200px] px-2 py-1 gap-y-1">
        <div>
            <Label for="topic">Topic</Label>
            <div class="flex gap-1">
                <Input
                    bind:value={topic}
                    id="topic"
                    onkeydown={(e) => {
                        if (e != null) {
                            if (e.key == "Enter") {
                                handleSave();
                            }
                        }
                    }}
                />
                <ButtonWithHelp
                    child={subscribeElement}
                    helpText={"Subscripe to Topic"}
                    topic={{}}
                />
            </div>
        </div>
        {#each $Topics as tp, i}
            <div
                class="my-1 odd:bg-gray-200 odd:bg-opacity-60 p-2 rounded-lg ring-1 ring-black flex justify-between"
            >
                <p class="">
                    {tp.topic}
                </p>
                <div class="flex gap-1">
                    <ButtonWithHelp
                        child={unsubscribeElement}
                        helpText={"Unsubscribe"}
                        topic={{topic:tp.topic}}
                    />
                    <ButtonWithHelp child={saveToLocal} helpText={"Save to browser"} topic={tp.topic} />
                </div>
            </div>
        {/each}
    </div>
</div>
