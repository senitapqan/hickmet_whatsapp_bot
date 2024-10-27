package models

type Session struct {
	Phone  string `json:"phone" db:"phone"`
	Status string `json:"status" db:"status"`
}
