package entity

import "github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/helper"

type User struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(50)"`
	Email string `json:"email" gorm:"type:varchar(50);unique"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Tasks []Task	`gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	helper.Model
}