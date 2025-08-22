package response

import "github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/database/entity"

type ListTask struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement"`
	NameTask string `json:"name_task" gorm:"type:varchar(100)"`
	Status string `json:"status"`
	TaskId uint `json:"task_id"`
	Model
}

func ParseAllLIstTask(data []entity.ListTask) (resp []ListTask) {
	for _, v := range data {
		resp = append(resp, ListTask{
			Id: v.Id,
			NameTask: v.NameTask,
			Status: v.Status,
			TaskId: v.TaskId,
			Model: Model{
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
		})
	}

	return resp
}