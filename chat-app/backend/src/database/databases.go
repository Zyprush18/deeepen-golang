package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100)"`
	Email string	`json:"email" gorm:"unique;type:varchar(100)"`
	Password string	`json:"password" gorm:"type:varchar(255)"`
	UUID string	`json:"uuid" gorm:"type:varchar(255)"`
}