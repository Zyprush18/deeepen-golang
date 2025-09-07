package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/auth"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/auth/repositoryauth"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/auth/servicesauth"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/message"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/middleware"
)

func RunApp() {
	initdb := database.ConnectDB()

	repoauth := repositoryauth.Connect(initdb)
	authsvc := servicesauth.NewService(&repoauth)
	authhandler := auth.NewHandle(authsvc)

	http.HandleFunc("/api/register", authhandler.Register)
	http.HandleFunc("/api/login", authhandler.Login)
	http.Handle("/api/profile", middleware.AuthMiddleware(http.HandlerFunc(authhandler.Profile)))

	h := message.NewHub()
	go h.Run()

	http.Handle("/ws/chat", middleware.MiddlewareWs(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message.ServeWs(h, w, r)
	})))

	fmt.Println("Websocket Running On Port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
