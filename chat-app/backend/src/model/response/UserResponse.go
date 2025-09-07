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
	Username string `json:"username"`
	Email string `json:"email"`
	Uuid string `json:"uuid"`
}