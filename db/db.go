package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserPostgres struct {
}

func Init() *gorm.DB {
	dbURL := "postgres://admin:123456@localhost:5432"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
