package model

import "time"

type ClientData struct {
	IdClient    int `gorm:"primaryKey;autoIncrement;"`
	IdUser      string
	Address     string
	AddressLong float64
	AddressLat  float64
	IsMale      bool
	Nik         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
