package repositoryfriend

import (
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/request"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/response"
	"gorm.io/gorm"
)

type Friend interface {
	ListAllFriend(id int) ([]response.FriendWithRole,error)
	CreateFriend(req *request.Friends) error
	UpdateFriend(req *request.Friends) error
	DeleteFriend(from_id,to_id int) error
}

type databases struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) databases  {
	return databases{db: d}
}

func (d *databases) ListAllFriend(id int) ([]response.FriendWithRole,error) {
	var modelfriend []response.FriendWithRole
	tx := d.db.Begin()

	if err:= tx.Table("friends").
	Debug().
	Distinct("friends.from_id,friends.to_user_id,friends.name_from, friends.name_to,u.uuid AS uuid,friends.status AS status, CASE WHEN friends.from_id = ? THEN 'sent' WHEN friends.to_user_id = ? THEN 'received' END AS ROLE",id,id).
	Joins("JOIN users u ON u.id = CASE WHEN friends.from_id = ? THEN friends.to_user_id ELSE friends.from_id END",id).
	Where("(friends.from_id = ? OR friends.to_user_id = ?) AND u.id <> ?",id,id,id).
	Scan(&modelfriend).Error;err != nil {
		tx.Rollback()
		return nil, err
	}

	return modelfriend,nil
}


func (d *databases) CreateFriend(req *request.Friends) error  {
	tx := d.db.Begin()

	var count int64
	if err:= tx.Table("friends").
	Where("(from_id = ? AND to_user_id=?) OR (from_id = ? AND to_user_id=?)", req.From_id,req.ToUser_id,req.ToUser_id,req.From_id).
	Count(&count).Error;err != nil {
		tx.Rollback()
		return err
	}

	if count > 0 {
		tx.Rollback()
		return gorm.ErrDuplicatedKey
	}

	if err:= tx.Table("friends").Create(req).Error;err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (d *databases) UpdateFriend(req *request.Friends) error  {
	tx := d.db.Begin()
	if err:= tx.Table("friends").
	Where("from_id=? AND (from_id = ? AND to_user_id=?)", req.From_id, req.From_id, req.ToUser_id).
	Updates(req).Error;err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (d *databases) DeleteFriend(from_id,to_id int) error  {
	tx := d.db.Begin()
	if err:= tx.Table("friends").Debug().
	Where("(from_id = ? AND to_user_id=?) OR (from_id=? AND to_user_id=?)", from_id,to_id, to_id,from_id).
	Delete(&database.Friend{}).Error;err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}