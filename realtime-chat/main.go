package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)




func main() {
	app := fiber.New()
	h := NewHub()
	go h.RunHub()

	// use the websocket
	app.Use("/ws", AllowWs)
	app.Use("/ws/chat", websocket.New(BidMessage(h)))




	log.Fatal(app.Listen(":3000"))
}