package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100)"`
	Email string	`json:"email" gorm:"unique;type:varchar(100)"`
	Password string	`json:"password" gorm:"type:varchar(255)"`
	UUID string	`json:"uuid" gorm:"type:varchar(255)"`
	Friends []Friend `gorm:"foreignKey:From_id;refrences:ID"`
}

type Friend struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100)"`
	From_id uint `json:"from_id"`
	Status string `json:"status" gorm:"default:pending"`
	ToUser_id uint `json:"to_user_id"`
	Users User `json:"users" gorm:"foreignKey:ToUser_id"`
}