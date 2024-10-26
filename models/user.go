package models

type User struct {
	FullName   string `json:"fullname" db:"fullname"`
	IIN        string `json:"iin" db:"iin"`
	Phone      string `json:"phone" db:"phone"`
	TotalMoney int    `json:"total_money" db:"total_money"`
}
