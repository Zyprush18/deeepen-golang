package response

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/database/entity"
)

type Task struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	SubTitle string `json:"sub_title"`
	Description string `json:"decription"`
	UserId uint `json:"user_id"`
	Model
}

func ParseAllTask(data []entity.Task) (resp []Task)  {
	for _, v := range data {
		resp = append(resp, Task{
			Id: v.Id,
			Title: v.Title,
			SubTitle: v.SubTitle,
			Description: v.Description,
			UserId: v.UserId,
			Model: Model{
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
		})
	}

	return resp
}