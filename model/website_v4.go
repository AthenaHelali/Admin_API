package model


type Website_v4 struct {
	ID          uint    `json:"website_id" gorm:"primarykey"`
	Type        string  `json:"ch_type"`
	Level       uint    `json:"level"`
	Fingerprint bool    `json:"fingerprint"`
	Brand       bool    `json:"brand"`
	Website     Website ``
}
