package config

import (
	"fmt"
	"log"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/database/migration"
	// "gorm.io/gorm"
)

func Config()   {
	e := newEnv()
	_,err := migration.NewConnection(e.DBName,e.DBHost,e.DBPort,e.DBUser,e.DBPass)
	if err != nil {
		fmt.Println("Failed Connect")
		log.Fatalln(err.Error())
	}
}
