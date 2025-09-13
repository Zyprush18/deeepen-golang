import { writable } from 'svelte/store';

const messageStore = writable('');
let socket:WebSocket | null= null;

// ws://localhost:3000/ws/chat?token=${token}&toUser=${uuid}
function ConnectWebSocket(token: any,uuid: any) {
	let conn = uuid === '' || uuid === null ? `ws://localhost:3000/ws/chat?token=${token}` : `ws://localhost:3000/ws/chat?token=${token}&toUser=${uuid}`

	if (socket) {
		socket.onclose = () => {
			openNewSocket(conn);
		};
		socket.close();
	} else {
		openNewSocket(conn);
	}
}


function openNewSocket(conn_url: any) {	
	socket = new WebSocket(conn_url);

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