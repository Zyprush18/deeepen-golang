package config

import (
	"fmt"
	"log"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/database/migration"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

func Config() *gorm.DB  {
	e := newEnv()
	DB,err := migration.NewConnection(e.DBName,e.DBHost,e.DBPort,e.DBUser,e.DBPass)
	if err != nil {
		fmt.Println("Failed Connect")
		log.Fatalln(err.Error())
	}

	return DB
}
