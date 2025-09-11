package response

// import "github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"

type Friend struct {
	From_id uint `json:"from_id"` 
	ToUser_id uint `json:"to_user_id"` 
	Name_From string `json:"name_from"`
	Name_To string `json:"name_to"`
	Uuid string `json:"uuid"`
	Status string `json:"status"`
}

type FriendWithRole struct {
	Friend
	Role  string `json:"role"`
}