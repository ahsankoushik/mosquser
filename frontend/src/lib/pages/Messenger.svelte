<script lang="ts">
    import mqtt from "mqtt";
    import { onMount } from "svelte";
    import type MosqClient from "../api/client.svelte";
    import { Topics } from "../api/stores";
    import MessengerSidebar from "../comps/MessengerSidebar.svelte";

    const { mc }:{mc:MosqClient} = $props();
    let Messages:Array<string> = $state([])

    // Define the MQTT broker URL and port number
    const brokerUrl = 'mqtt://localhost:9001';

    // Define the MQTT client options
    const clientOptions = {
    username: "ahsankoushik@gmail.com",
    password: 'koushik123',
    clientId: 'mqttjs_' + Math.random().toString(16).substring(2, 10),
    clean: true
    };
    let client: mqtt.MqttClient;
    onMount(()=>{
        console.log("connecting to mqtt broker.....");
        // Create a new MQTT client instance
        client =  mqtt.connect(brokerUrl, clientOptions);

        // Define a callback function to handle the connection event
        client.on('connect', function () {
        
        });

        client.on("error",(e)=>{
        console.log(e);
        })

        // Define a callback function to handle the message event
        client.on('message', function (topic, message) {
            console.log(topic,message.toString());

        });
    })
    Topics.subscribe((t)=>{
        
    })
</script>
<div class="flex w-full h-full">
    <MessengerSidebar {mc}/>

    <div class="w-[70vw] min-w-[400px] border flex-shrink-0">
    </div>
</div>