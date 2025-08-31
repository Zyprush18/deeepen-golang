package message

import (
	"log"
	"net/http"
	"time"

	"github.com/Zyprush18/deeepen-golang/chat-app/src/helper"
	"github.com/Zyprush18/deeepen-golang/chat-app/src/model/response"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	// check who makes request
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:3000"
	},
}

type Client struct {
	hub *Hub

	conn *websocket.Conn

	send chan *response.MessageResponse
}

func (c *Client) readMessage(n string) {
	defer func() {
		c.hub.RemoveClient <- c
		c.conn.Close()
	}()

	// set read deadline
	c.conn.SetReadDeadline(time.Now().Add(time.Minute))
	c.conn.SetPongHandler(func(appData string) error {
		return c.conn.SetReadDeadline(time.Now().Add(time.Minute))
	})
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error: ", err)
			}
			break
		}

		c.hub.Broadcast <- &response.MessageResponse{Name: n, Message: string(msg)}
	}
}

func (c *Client) writeMessage() {
	ticker := time.NewTicker(60 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteJSON(msg); err != nil {
				log.Println("Failed Write: ", err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWs(h *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed Upgrade: ", err)
		return
	}

	client := &Client{hub: h, conn: conn, send: make(chan *response.MessageResponse)}
	client.hub.RegisterClient <- client

	name := r.Context().Value(helper.Name).(string)

	go client.writeMessage()
	go client.readMessage(name)
}
