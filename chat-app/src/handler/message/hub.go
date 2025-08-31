package message

import (
	"github.com/Zyprush18/deeepen-golang/chat-app/src/model/response"
)

type Hub struct {
	Client         map[*Client]bool
	RegisterClient chan *Client
	RemoveClient   chan *Client
	Broadcast      chan *response.MessageResponse
}

func NewHub() *Hub {
	return &Hub{
		Client:         make(map[*Client]bool),
		RegisterClient: make(chan *Client),
		RemoveClient:   make(chan *Client),
		Broadcast:      make(chan *response.MessageResponse),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.RegisterClient:
			h.Client[client] = true
		case client := <-h.RemoveClient:
			if h.Client[client] {
				delete(h.Client, client)
				close(client.send)
			}
		case msg := <-h.Broadcast:
			for clients := range h.Client {
				select {
				case clients.send <- msg:
				default:
					close(clients.send)
					delete(h.Client, clients)
				}
			}
		}
	}
}
