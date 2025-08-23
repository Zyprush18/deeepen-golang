package repository

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/response"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(req *request.User) error
	Login(email string) (*response.User,error)
}

type database struct {
	db *gorm.DB
}

func NewConnection(d *gorm.DB) database {
	return database{db: d}
}

func (d *database)  Register(req *request.User) error {
	if err:=d.db.Table("users").Create(req).Error;err != nil {
		return err
	}
	return nil
}

func (d *database) Login(email string) (*response.User,error) {
	var modeluser response.User
	if err:= d.db.Table("users").Where("email =?", email).First(&modeluser).Error;err != nil {
		return nil, err
	}

	return &modeluser,nil
}