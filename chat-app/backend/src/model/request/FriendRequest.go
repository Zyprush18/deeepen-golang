package request

import (
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"
)

type Friends struct {
	database.Model
	Name_From string `json:"name_from"`
	Name_To string `json:"name_to"`
	From_id uint `json:"from_id"`
	Status string `json:"status"` // field for update friends in service or logic business
	ToUser_id uint `json:"to_user_id"`
}