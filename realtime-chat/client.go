package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func AllowWs(c *fiber.Ctx) error  {
	// upgrade ke connection websocket
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

func BidMessage(h *Hub) func(conn *websocket.Conn)   {
	return func(conn *websocket.Conn) {
		defer func ()  {
			h.removeclient <- conn	
			conn.Close()
		}()

		name := conn.Query("name","Agus")
		h.register <- conn

		for {
			mt,msg,err := conn.ReadMessage()
			if err != nil {
				log.Println("Read: ", err)
				return
			}

			// mengecek apakh type message nya sama atau nggak
			if mt == websocket.TextMessage {
				h.broadcast <- Message{Name: name, Message: string(msg)}
			}
		}
	}
}