package migration

import (
	"fmt"
	"log"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/database/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewConnection(db,host,port,username,password string) (DB *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host,username,password,db,port)
	DB,err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err := DB.AutoMigrate(&entity.User{},&entity.Task{},&entity.ListTask{});err!= nil{
		fmt.Println("Failed Migrate")
		log.Fatalln(err.Error())
	} 

	return DB,err

}