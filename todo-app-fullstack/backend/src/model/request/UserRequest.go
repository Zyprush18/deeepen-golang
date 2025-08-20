package request

import "github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/helper"

type User struct {
	Username string `json:"username" binding:"required,min:3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min:6,max:12"`
	helper.Model
}