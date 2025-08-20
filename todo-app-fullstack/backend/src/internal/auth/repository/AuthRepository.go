package repository

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(req *request.User) error
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