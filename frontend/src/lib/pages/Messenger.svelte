<script lang="ts">
    import mqtt from "mqtt";
    import { onMount } from "svelte";
    import type MosqClient from "../api/client.svelte";

    const { mc }:{mc:MosqClient} = $props();

    // Define the MQTT broker URL and port number
    const brokerUrl = 'mqtt://localhost:9001';

    // Define the MQTT client options
    const clientOptions = {
    username: "ahsankoushik@gmail.com",
    password: 'koushik123',
    clientId: 'mqttjs_' + Math.random().toString(16).substring(2, 10),
    clean: true
    };
    onMount(()=>{
        console.log("connecting to mqtt broker.....");
        // Create a new MQTT client instance
        const client =  mqtt.connect(brokerUrl, clientOptions);

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

        // Subscribe to a topic
        client.subscribe('#');

        return client;
    })
</script>