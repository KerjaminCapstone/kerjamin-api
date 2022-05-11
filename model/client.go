package model

import "time"

type ClientData struct {
	IdClient    int `gorm:"primaryKey;autoIncrement;"`
	IdUser      string
	HousePict   string
	PhoneNumber string
	Address     string
	AddressLong float64
	AddressLat  float64
	IsMale      bool
	Dob         time.Time
	Nik         string
	ProfilePict string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
