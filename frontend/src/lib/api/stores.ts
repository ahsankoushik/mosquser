import { writable } from "svelte/store";


export interface Topic{
    topic: string,
    connected: boolean,
    saved: boolean
}

export const Topics = writable<Array<Topic>>([])
