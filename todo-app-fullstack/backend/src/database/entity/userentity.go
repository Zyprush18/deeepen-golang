package entity

type User struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(50)"`
	Email string `json:"email" gorm:"type:varchar(50);unique"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Tasks []Task	`gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Model
}