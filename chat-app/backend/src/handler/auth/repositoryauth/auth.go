package repositoryauth

import (
	// "fmt"

	"fmt"

	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/request"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/response"

	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(req *request.Register) error
	Login(email string) (*response.Auth, error)
	Profile(email string,id int) ([]database.User, error)
}

type databases struct {
	db *gorm.DB
}

func Connect(d *gorm.DB) databases {
	return databases{db: d}
}

func (d *databases) Register(req *request.Register) error {
	if err := d.db.Table("users").Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (d *databases) Login(email string) (*response.Auth, error) {
	var modelauth response.Auth
	if err := d.db.Table("users").Where("email = ?", email).First(&modelauth).Error; err != nil {
		return nil, err
	}

	return &modelauth, nil
}

func (d *databases) Profile(email string,id int) ([]database.User, error) {
	var modelauth database.User
	var modelauthtest []database.User
	if err := d.db.Model(&modelauth).Preload("Friends.Users").Where("email = ?", email).First(&modelauth).Error; err != nil {
		return nil, err
	}

	if err:= d.db.Model(&modelauth).Preload("Friends.Users").Debug().Joins("JOIN friends ON friends.from_id = users.id OR friends.to_user_id = users.id").Where("email = ? AND (friends.from_id = ? OR friends.to_user_id = ?)", email, id, id).Scan(&modelauthtest).Error;err != nil {
		return nil, err
	}
	fmt.Println(modelauthtest)

	// fmt.Println(modelauthtest)

	// return &response.Profile{
	// 	Id: modelauth.ID,
	// 	Username: modelauth.Username,
	// 	Email: modelauth.Email,
	// 	Uuid: modelauth.UUID,
	// 	Friends: response.ParseFriend(modelauth.Friends),
	// }, nil

	return modelauthtest,nil
}