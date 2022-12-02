package model

import "time"

type Website struct {
	ID          uint      `json:"id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	User_id     uint      `json:"user_id"`
	Site_key    string    `json:"site_key"`
	Secret_key  string    `json:"secret_key"`
	Lable       string    `json:"lable"`
	Alert       bool      `json:"alert"`
	Subdomain   bool      `json:"subdomain"`
	Version     uint      `json:"version"`
	Alert_limit uint      `json:"alert_limit"`
}
