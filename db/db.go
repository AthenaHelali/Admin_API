package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserPostgres struct {
}

func Init() *gorm.DB {
	time.Sleep(5 * time.Second)
	dbURL := "postgres://admin:123456@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
