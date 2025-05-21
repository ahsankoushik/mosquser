import { writable } from "svelte/store";


interface Topic{
    topic: string,
    connected: boolean
}

export const Topics = writable<Array<Topic>>([])