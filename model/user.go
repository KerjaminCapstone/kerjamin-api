package model

import "time"

type User struct {
	IdUser       string `gorm:"primaryKey"`
	Name         string `validate:"required"`
	Username     string `validate:"required"`
	Email        string `validate:"required"`
	Password     string `validate:"required,gte=5,lte=30"`
	RefreshToken string
	IdToken      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
