package response

import (
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"
)

type Auth struct {
	database.User
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Uuid     string `json:"uuid"`
}

type Profile struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Uuid     string `json:"uuid"`
}

func ParseProfile(data *database.User) *Profile {
	return &Profile{
		Id:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Uuid:     data.UUID,
	}
}
