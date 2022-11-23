package model

type User struct {
	ID                     uint64 `json:"id" gorm:"primaryKey"`
	Created_at             string `json:"created at"`
	Updated_at             string `json:"updated_at"`
	Ouath_id               string `json:"ouath_id"`
	Email                  string `json:"email"`
	Name                   string `json:"name"`
	Password               string `json:"password"`
	Phone                  string `json:"phone"`
	Company_name           string `json:"company_name"`
	Job_title              string `json:"job_title"`
	Active                 bool   `json:"active"`
	Subscribe_news         bool   `json:"subscribe_news"`
	Subscribe_notification bool   `json:"subscribe_notification"`
}
