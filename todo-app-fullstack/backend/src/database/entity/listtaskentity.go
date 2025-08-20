package entity

import "github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/helper"

type ListTask struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement"`
	NameTask string `json:"name_task" gorm:"type:varchar(100)"`
	Status string `json:"status"`
	TaskId uint `json:"task_id"`
	Tasks Task `gorm:"foreignKey:TaskId"`
	helper.Model
}