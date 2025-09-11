package request

import "github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"

type Register struct {
	database.Model
	Username string	`json:"username"`
	Email string	`json:"email"`
	Password string	`json:"password"`
	Uuid string	`json:"uuid"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}