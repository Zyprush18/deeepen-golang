import { writable } from 'svelte/store';

const messageStore = writable('');
let socket:WebSocket;

function ConnectWebsocket(token: any) {
	socket = new WebSocket(`ws://localhost:3000/ws/chat?token=${token}`);

	socket.addEventListener('open', function (event) {
		console.log("it's Open");
	});

	socket.addEventListener('message', (event) => {
		messageStore.set(event.data)
	});
}

function sendMessage(message:string) {
    if (socket.readyState <= 1) {
        socket.send(message)
    }
}
export default {
    connect: ConnectWebsocket,
    subscribe: messageStore.subscribe,
    send:sendMessage
}