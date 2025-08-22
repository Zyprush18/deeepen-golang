package request

type Task struct {
	Title string `json:"title" binding:"required,min=3"`
	SubTitle string `json:"sub_title"`
	Description string `json:"decription"`
	UserId uint `json:"user_id"`
	Model
}