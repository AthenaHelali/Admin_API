package model

type Website_v1 struct {
	Website_id  uint   `json:"website_id" gorm:"ref"`
	Type        string `json:"ch_type"`
	Level       uint   `json:"level"`
	Fingerprint bool   `json:"fingerprint"`
	Brand       bool   `json:"brand"`
}
