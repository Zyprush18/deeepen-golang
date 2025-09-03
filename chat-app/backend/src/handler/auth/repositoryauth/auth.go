package repositoryauth

import (
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/request"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/response"

	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(req *request.Register) error
	Login(email string) (*response.Auth, error)
}

type database struct {
	db *gorm.DB
}

func Connect(d *gorm.DB) database {
	return database{db: d}
}

func (d *database) Register(req *request.Register) error {
	if err := d.db.Table("users").Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (d *database) Login(email string) (*response.Auth, error) {
	var modelauth response.Auth
	if err := d.db.Table("users").Where("email = ?", email).First(&modelauth).Error; err != nil {
		return nil, err
	}

	return &modelauth, nil
}
