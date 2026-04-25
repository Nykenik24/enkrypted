import { writable } from 'svelte/store';
import { browser } from '$app/environment';

export type WSMessage = {
    kind: string;
    data: any;
    auth: any;
};

const { subscribe, set, update } = writable({
    connected: false,
    messages: [] as WSMessage[]
});

let socket: WebSocket | null = null;

export function connect(wsUrl: string) {
    if (!browser) return;

    if (socket) return;

    socket = new WebSocket(wsUrl);

    socket.onopen = () => {
        update((state) => ({ ...state, connected: true }));
    };

    socket.onmessage = (event) => {
        const data = JSON.parse(event.data);
        update((state) => ({
            ...state,
            messages: [...state.messages, data]
        }));
    };

    socket.onclose = () => {
        set({ connected: false, messages: [] });
        socket = null;
    };
}

export function send(data: any) {
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify(data));
    }
}

export function disconnect() {
    if (socket) {
        socket.close();
        socket = null;
    }
}

export const wsStore = { subscribe };
