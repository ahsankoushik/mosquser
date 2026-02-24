<script lang="ts">
    import mqtt from "mqtt";
    import { onMount } from "svelte";
    import type MosqClient from "../api/client.svelte";
    import MessengerSidebar from "../comps/MessengerSidebar.svelte";
    import Loading from "../comps/Loading.svelte";
    import Label from "../components/ui/label/label.svelte";
    import Button from "../components/ui/button/button.svelte";
    import { Send } from "@lucide/svelte";
    import toast from "svelte-5-french-toast";

    const { mc }:{mc:MosqClient} = $props();
    let mqttConnected = $state(false);
    let messages:Array<{topic:string, msg:string,sent:boolean}> = $state([])
    let bottomElement : HTMLDivElement;
    let sendTopic = $state("");
    let sendMsg = $state("");

    // Define the MQTT broker URL and port number


    // Define the MQTT client options
    const clientOptions = {
        username: mc.tokenData.username,
        password: mc.token,
        clientId: 'mqttjs_' + Math.random().toString(16).substring(2, 10),
        clean: true
    };
    let client: mqtt.MqttClient|undefined;
    onMount(async()=>{
        setTimeout(() => {
            if(!mqttConnected){
                mqttConnected = true;
                client = undefined;
                toast.error("Failed to connect to MQTT broker. Please check your connection settings.")
            }
        }, 5000);
        const res = await mc.getKeys(["broker_url"]);
        if(res.data.length == 0 ){
            return
        }
        let brokerUrl = "mqtt://localhost:90081";
        for(let i = 0; i < res.data.length; i++){
            if(res.data[i].key = "broker_url"){
                brokerUrl = res.data[i].value;
            }
        }
        console.log("connecting to mqtt broker.....");
        // Create a new MQTT client instance
        client =  mqtt.connect(brokerUrl, clientOptions);

        // Define a callback function to handle the connection event
        client.on('connect', function () {
            mqttConnected = true;
        });

        client.on("error",(e)=>{
            console.log(e);
        })

        // Define a callback function to handle the message event
        client.on('message', function (topic, message) {
            // console.log(topic,message.toString());
            messages.push({
                topic,
                msg:message.toString(),
                sent:false
            });
        });
    })
    const subscribe = (topic:string)=>{
        if(client == undefined){
            return
        }
        client.subscribe(topic)
    }
    const unsubscribe = (topic:string) =>{ 
        if(client == undefined){
            return
        }
        client.unsubscribe(topic)
    }
    const scrollToBottom = ()=>{
        if(bottomElement == undefined){
            return
        }
        bottomElement.scrollIntoView({behavior:"smooth"})
    }
    $effect(()=>{
        messages.length;
        scrollToBottom();
    })
    
</script>
<svelte:head>
    <title>Messenger | Mosquser</title>
</svelte:head>
{#if !mqttConnected}
    <Loading text="Connecting to Mosquitto."/>
{/if}
<div class="flex w-full h-full">
    <MessengerSidebar 
        {mc}
        {subscribe}
        {unsubscribe}
    />
    <div class="w-[70vw] min-w-[400px] border  relative overflow-x-scroll"
    >
        {#each messages as msg}
            <div
                class="flex m-2 {msg.sent ? "justify-end " : "justify-start "}"
            >
                <div
                    class="p-3 break-all min-w-[100px]  max-w-[50%] rounded-lg {msg.sent ? " bg-green-400" : " bg-gray-200"}"
                >
                    <h3><span class="font-semibold mr-2">Topic: </span>{msg.topic}</h3>
                    <p>{msg.msg}</p>
                </div>
            </div>
        {/each}
        <div bind:this={bottomElement} class="h-20 "></div>
        <div class="fixed bottom-0 z-10 w-full border-t flex gap-2 p-1 bg-white overflow-x-scroll">
            <div class="w-[20vw] min-w-[120px]">
                <Label for="topic-send">Topic</Label>
                <input id="topic-send"
                    bind:value={sendTopic}
                    placeholder="Enter topic here."
                    class="w-full border rounded-md p-1 m-1 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring"
                />
            </div>
            <div class="w-[44vw] min-w-[250px]">
                <textarea
                    id=""
                    bind:value={sendMsg}
                    placeholder="Enter message here."
                    class="w-full h-full resize-none px-1 border rounded-md focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring"
                ></textarea>
            </div>
            <div class="">
                <Button
                    onclick={()=>{
                        if(sendTopic == "" || sendMsg == ""){
                            return
                        }
                        if(client == undefined){
                            return
                        }
                        console.log(client.publish(sendTopic,sendMsg));
                        messages.push({
                            topic:sendTopic,
                            msg:sendMsg,
                            sent:true
                        })
                    }}
                ><Send/></Button>
            </div>
        </div>
    </div>
    
</div>
