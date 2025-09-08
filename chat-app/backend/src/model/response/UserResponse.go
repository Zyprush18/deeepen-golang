package response

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Username string `json:"username"`
	Email string	`json:"email"`
	Password string	`json:"password"`
	Uuid string	`json:"uuid"`
}

type Profile struct {
	Id  uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Uuid string `json:"uuid"`
	Friends []Friend `json:"friend"`
}