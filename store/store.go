package store

import (
	"context"
	"log"

	"gorm.io/gorm"
	"main.go/model"
)

type UserPostgre struct {
	db *gorm.DB
}

func NewUserPostgres(database *gorm.DB) *UserPostgre {
	return &UserPostgre{db: database}
}
func (store *UserPostgre) Save(ctx context.Context, m *model.User) error {
	if err := store.db.Create(m).Error; err != nil {
		log.Printf("user creation on Postgre failed: %v", err)
		return err
	}
	return nil
}
