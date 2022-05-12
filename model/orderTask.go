package model

import "time"

type OrderTask struct {
	IdTask     string `gorm:"primaryKey"`
	IdOrder    string
	TaskDesc   string
	TaskStatus bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
