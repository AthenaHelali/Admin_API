package model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID                      uint      `json:"id" gorm:"primaryKey" gorm:"foreignKey"`
	Created_at              time.Time `json:"created_at"`
	Updated_at              time.Time `json:"updated_at"`
	Oauth_id                string    `json:"oauth_id"`
	Email                   string    `json:"email"`
	Name                    string    `json:"name"`
	Password                string    `json:"password" validate:"min=8,max=16"`
	Phone                   string    `json:"phone"`
	Company_name            string    `json:"company_name"`
	Job_title               string    `json:"job_title"`
	Active                  bool      `json:"active"`
	Subscribe_news          bool      `json:"subscribe_news"`
	Subscribe_notifications bool      `json:"subscribe_notifications"`
}

func HashPassword(pass string) (string, error) {
	if len(pass) == 0 {
		return "", errors.New("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash), err
}
