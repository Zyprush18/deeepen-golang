package message

import (
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/response"
)

type Hub struct {
	Client         map[string]map[*Client]bool
	RegisterClient chan *Client
	RemoveClient   chan *Client
	Broadcast      chan *response.MessageResponse
}

func NewHub() *Hub {
	return &Hub{
		Client:         make(map[string]map[*Client]bool),
		RegisterClient: make(chan *Client),
		RemoveClient:   make(chan *Client),
		Broadcast:      make(chan *response.MessageResponse),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.RegisterClient:
			// di cek apakah user udah pernah connect atau belum
			if _, ok := h.Client[client.IdUser]; !ok {
				h.Client[client.IdUser] = make(map[*Client]bool)
			}
			h.Client[client.IdUser][client] = true

		case client := <-h.RemoveClient:
			// cek apakah client dengan user id tertentu ada atau nggak
			if cons, ok := h.Client[client.IdUser]; ok {
				// mengecek koneksi dari client tersebut
				if _, ok := cons[client]; ok {
					delete(cons, client)
					close(client.send)
				}

				// mengecek apakah koneksi client nya itu masih ada atau nggak
				if len(cons) == 0 {
					delete(h.Client, client.IdUser)
					close(client.send)
				}
			}
		case msg := <-h.Broadcast:
			// private chat
			if msg.To != "" {
				// mengecek apakah si user ini koneksi nya ada atau nggak
				if conn, ok := h.Client[msg.From]; ok {
					for clients := range conn {
						select {
						case clients.send <- msg:
						default:
							delete(conn, clients)
							close(clients.send)
						}
					}
				}

				// mengecek apakah si user yg di tuju ini koneksi nya ada atau nggak
				if conn, ok := h.Client[msg.To]; ok {
					for clients := range conn {
						select {
						case clients.send <- msg:
						default:
							delete(conn, clients)
							close(clients.send)
						}
					}
				}
			} else {
				// public chat
				for _, clients := range h.Client {
					for conns := range clients {
						select {
						case conns.send <- msg:
						default:
							delete(clients, conns)
							close(conns.send)
						}
					}
				}
			}
		}
	}
}
