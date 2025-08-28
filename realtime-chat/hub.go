package main

import "github.com/gofiber/contrib/websocket"

// struct for response
type Message struct {
	Name string
	Message string
}

type Hub struct {
	// for added client to websocket
	clients map[*websocket.Conn]bool
	// register client
	register chan *websocket.Conn
	// remove client
	removeclient chan *websocket.Conn
	// broadcast message from client
	broadcast chan Message
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]bool),
		register: make(chan *websocket.Conn),
		removeclient: make(chan *websocket.Conn),
		broadcast: make(chan Message),
	}
}

func (h *Hub) RunHub()  {
	for {
		select{
		case conn := <- h.register:
			h.clients[conn] = true
		case conn := <- h.removeclient:
			delete(h.clients, conn)
		case msg:=  <- h.broadcast:
			for  client := range h.clients {
				_ = client.WriteJSON(msg)
			}
		}
	}
}
