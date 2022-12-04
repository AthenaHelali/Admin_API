package model

type Website_v1 struct {
	ID          uint   `json:"website_id"`
	Type        string `json:"ch_type"`
	Level       uint   `json:"level"`
	Fingerprint bool   `json:"fingerprint"`
	Brand       bool   `json:"brand"`
}
