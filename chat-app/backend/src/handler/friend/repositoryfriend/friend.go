package repositoryfriend

// import (
// 	"fmt"

// 	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"
// 	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/response"
// 	"gorm.io/gorm"
// )

// type Friend interface {
// 	GetFriend(id int)(*response.UserFriend,error)
// }

// type databases struct {
// 	db *gorm.DB
// }

// func ConnectDb(d *gorm.DB) databases  {
// 	return databases{db: d}
// }

// func (d *databases) GetFriend(id int)(*response.UserFriend,error) {
// 	var modelfriend database.User

// 	if err:= d.db.Debug().Model(&modelfriend).Preload("Friends").Joins("JOIN friends as f ON f.to_user_id = users.id").Where("f.from_id = ?", id).First(&modelfriend).Error;err != nil {
// 		return nil, err
// 	}

// 	fmt.Println(modelfriend)

// 	return &response.UserFriend{
// 		Username: modelfriend.Username,
// 		Email: modelfriend.Email,
// 		Uuid: modelfriend.UUID,
// 		Friends: response.ParseFriend(modelfriend.Friends),
// 	},nil
// }