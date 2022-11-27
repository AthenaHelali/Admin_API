package model

type website struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Type        string `json:"ch_type"`
	Level       uint   `json:"level"`
	Fingerprint bool   `json:"fingerprint"`
	Brand       bool   `json:"brand"`
}
