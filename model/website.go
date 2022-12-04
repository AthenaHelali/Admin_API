package model

import "time"

type Website struct {
	ID          uint       `json:"id"`
	Created_at  time.Time  `json:"created_at"`
	Updated_at  time.Time  `json:"updated_at"`
	User_id     uint       `json:"user_id"`
	Site_key    string     `json:"site_key"`
	Secret_key  string     `json:"secret_key"`
	Label       string     `json:"label"`
	Alert       bool       `json:"alert"`
	Subdomain   bool       `json:"subdomain"`
	Version     uint       `json:"version"`
	Alert_limit uint       `json:"alert_limit"`
	Website_v1  Website_v1 `gorm:"foreignKey:ID"`
}
