package request

type Friends struct {
	Name string `json:"name"`
	From_id uint `json:"from_id"`
	Status string `json:"status"`
	ToUser_id uint `json:"to_user_id"`
}