<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import type MosqClient from "../api/client.svelte";
    import { Topics, type Topic } from "../api/stores";
    import { get } from "svelte/store";
    import Input from "../components/ui/input/input.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { BadgePlus, Save, Trash2, Unplug } from "@lucide/svelte";
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

    onMount(() => {
        const topicsStr = localStorage.getItem("mqtt_topics");
        if (topicsStr != null) {
            const topics = JSON.parse(topicsStr);
            let topicsArr: Array<Topic> = [];
            for (let i = 0; i < topics.length; i++) {
                topicsArr.push({
                    topic: topics[i],
                    connected: false,
                    saved: true,
                });
            }
            Topics.set(topicsArr);
        }
    });
    const handleSubscribe = () => {
        if (topic == "") {
            return;
        }
        subscribe(topic);
        Topics.update((t) => {
            t.push({
                topic,
                connected: true,
                saved: false,
            });
            return t;
        });
        topic = "";
    };
    const handleSavedSubscribe = (topic: string) => {
        subscribe(topic);
        Topics.update((t) => {
            for (let i = 0; i < t.length; i++) {
                if (t[i].topic == topic) {
                    t[i].connected = true;
                }
            }
            return t;
        });
    };
    const handleUnsubscribe = (topic: string) => {
        if (topic == "") {
            return;
        }
        unsubscribe(topic);
        Topics.update((t) => {
            t = t.filter((e) => {
                if (e.topic == topic) {
                    if (e.saved) {
                        e.connected = false;
                        return true;
                    }
                    return false;
                }
                return true;
            });
            return t;
        });
    };
    const handleSaveToLocal = (topic: string) => {
        if (topic == "") {
            return;
        }
        let topics = localStorage.getItem("mqtt_topics");
        let topicsArr = [];
        if (topics != null) {
            topicsArr = JSON.parse(topics);
            console.log(topics);
        }
        topicsArr.push(topic);
        Topics.update((t) => {
            for (let i = 0; i < t.length; i++) {
                if (t[i].topic == topic) {
                    t[i].saved = true;
                }
            }
            return t;
        });
        localStorage.setItem(
            "mqtt_topics",
            JSON.stringify([...new Set(topicsArr)]),
        );
    };
    const handleDeleteFromLocal = (topic: string) => {
        if (topic == "") {
            return;
        }
        let topicsArr = JSON.parse(
            localStorage.getItem("mqtt_topics") || "",
        ) as Array<string>;
        topicsArr = topicsArr.filter((e) => {
            return e != topic;
        });
        Topics.update((t) => {
            t = t.filter((a) => {
                if (a.topic != topic) {
                    if (a.connected) {
                        a.saved = false;
                        console.log(a.saved);
                        return true;
                    }
                    return false;
                }
                return true;
            });
            return t;
        });
        console.log(get(Topics))
        localStorage.setItem("mqtt_topics", JSON.stringify(topicsArr));
    };
</script>

{#snippet subscribeElement()}
    <Button onclick={handleSubscribe}>
        <BadgePlus />
    </Button>
{/snippet}
{#snippet subscribeSaved({ topic }: { topic: string })}
    <button
        class="hover:bg-gray-300 p-1 rounded-lg"
        onclick={() => {
            handleSavedSubscribe(topic);
        }}><BadgePlus /></button
    >
{/snippet}
{#snippet saveToLocal({ topic }: { topic: string })}
    <button
        class="hover:bg-gray-300 p-1 rounded-lg"
        onclick={() => {
            handleSaveToLocal(topic);
        }}><Save /></button
    >
{/snippet}
{#snippet deleteFromLocal({ topic }: { topic: string })}
    <button
        class="hover:bg-gray-300 p-1 rounded-lg"
        onclick={() => {
            handleDeleteFromLocal(topic);
        }}><Trash2 /></button
    >
{/snippet}
{#snippet unsubscribeElement({ topic }: { topic: string })}
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
                                handleSubscribe();
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
                    {#if tp.connected}
                        <ButtonWithHelp
                            child={unsubscribeElement}
                            helpText={"Unsubscribe"}
                            topic={{ topic: tp.topic }}
                        />
                    {:else}
                        <ButtonWithHelp
                            child={subscribeSaved}
                            helpText={"Subscribe"}
                            topic={{ topic: tp.topic }}
                        />
                    {/if}
                    {#if tp.saved}
                        <ButtonWithHelp
                            child={deleteFromLocal}
                            helpText={"Delete From Browser"}
                            topic={{ topic: tp.topic }}
                        />
                    {:else}
                        <ButtonWithHelp
                            child={saveToLocal}
                            helpText={"Save to browser"}
                            topic={{ topic: tp.topic }}
                        />
                    {/if}
                </div>
            </div>
        {/each}
    </div>
</div>
