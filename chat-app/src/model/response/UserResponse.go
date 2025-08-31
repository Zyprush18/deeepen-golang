package response

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Username string `json:"username"`
	Email string	`json:"email"`
	Password string	`json:"password"`
}