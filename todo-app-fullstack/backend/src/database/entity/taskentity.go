package entity


type Task struct {
	Id uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Title string `json:"title" gorm:"varchar(100)"`
	SubTitle string `json:"sub_title" gorm:"varchar(100)"`
	Description string `json:"decription" gorm:"varchar(200);default:null"`
	UserId uint `json:"user_id"`
	Users User `gorm:"foreignKey:UserId"`
	ListTasks []ListTask `gorm:"foreignKey:TaskId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Model
}