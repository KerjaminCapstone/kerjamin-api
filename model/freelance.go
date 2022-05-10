package model

import (
	"time"
)

type FreelanceData struct {
	IdFreelance int `gorm:"primaryKey;autoIncrement;"`
	IdUser      string
	IsTrainee   bool
	Rating      float64
	JobDone     int
	DateJoin    time.Time
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
