package model

import "time"

type OrderPayment struct {
	IdPayment int `gorm:"primaryKey;autoIncrement;"`
	IdOrder   string
	Value     int
	IdMethod  int
	IsPaid    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
