package database

type User struct {
	Model
	Username string `json:"username" gorm:"type:varchar(100)"`
	Email string	`json:"email" gorm:"unique;type:varchar(100)"`
	Password string	`json:"password" gorm:"type:varchar(255)"`
	UUID string	`json:"uuid" gorm:"type:varchar(255)"`
	FriendsFrom []Friend `gorm:"foreignKey:From_id;references:ID"`
	FriendsTo []Friend `gorm:"foreignKey:ToUser_id;references:ID"`
}

type Friend struct {
	Model
	Name_From string `json:"name_from" gorm:"type:varchar(100)"`
	Name_To string `json:"name_to" gorm:"type:varchar(100)"`
	From_id uint `json:"from_id"`
	Status string `json:"status" gorm:"default:pending"`
	ToUser_id uint `json:"to_user_id"`
	UsersFrom User `json:"users_from" gorm:"foreignKey:From_id"`
	UsersTo User `json:"users_to" gorm:"foreignKey:ToUser_id"`
}