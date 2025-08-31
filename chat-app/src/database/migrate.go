package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=123 dbname=chatrealtime port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed Connect",err)
		return nil
	}

	// migrate
	if err:= db.AutoMigrate(&User{});err != nil {
		log.Fatal("Failed Migrate: ", err)
		return nil
	}

	return db
}
