package main

import (
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
)

func BidMessage(h *Hub) func(conn *websocket.Conn) {
	return func(conn *websocket.Conn) {
		ticker := time.NewTicker(30 * time.Second)

		defer func() {
			ticker.Stop()
			h.removeclient <- conn
			conn.Close()
		}()

		name := conn.Query("name", "Agus")
		h.register <- conn

		conn.SetPongHandler(func(appData string) error {
			conn.SetReadDeadline(time.Now().Add(time.Minute))
			return nil
		})

		go func() {
			<-ticker.C

			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Ping Error: ", err)
				return
			}
		}()

		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read: ", err)
				return
			}

			// mengecek apakh type message nya sama atau nggak
			if mt == websocket.TextMessage {
				conn.SetWriteDeadline(time.Now().Add(time.Minute))
				h.broadcast <- Message{Name: name, Message: string(msg)}
			}
		}
	}
}
