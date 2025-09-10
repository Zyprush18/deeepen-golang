import { writable } from 'svelte/store';

const messageStore = writable('');
let socket:WebSocket | null= null;

function ConnectWebSocket(token: any,uuid: any) {
	if (socket) {
		socket.onclose = () => {
			openNewSocket(token, uuid);
		};
		socket.close();
	} else {
		openNewSocket(token, uuid);
	}
}


function openNewSocket(token: any,uuid: any) {	
	socket = new WebSocket(`ws://localhost:3000/ws/chat?token=${token}&toUser=${uuid}`);

	socket.addEventListener('open', function (event) {
		console.log("it's Open");
	});

	socket.addEventListener('message', (event) => {
		messageStore.set(event.data)
	});
}

async function sendMessage(message:string) {
    if (socket && socket.readyState === socket.OPEN) {
        socket.send(message)
    }
}
export default {
    connect: ConnectWebSocket,
    subscribe: messageStore.subscribe,
    send:sendMessage
}