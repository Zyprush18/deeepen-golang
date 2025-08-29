package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)


func AllowWs(c *fiber.Ctx) error  {
	// upgrade to connection websocket
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

func main() {
	app := fiber.New()
	h := NewHub()
	go h.RunHub()

	// use the websocket
	app.Use("/ws", AllowWs)
	app.Get("/ws/chat", websocket.New(BidMessage(h)))

	log.Fatal(app.Listen(":3000"))
}