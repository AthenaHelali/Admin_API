package model

import (
	"time"

	"gorm.io/gorm"
)

type User_plan struct {
	ID         uint           `json:"id" gorm:"primaryKay"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
	Plan_id    uint           `json:"plan_id" gorm:"primaryKay"`
	User_id    uint           `json:"user_id" gorm:"primaryKay"`
	Ex_time    time.Time      `json:"ex_time"`
}
