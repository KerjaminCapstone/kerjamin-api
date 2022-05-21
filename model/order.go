package model

import "time"

type Order struct {
	IdOrder        string `gorm:"primaryKey"`
	IdClient       int64
	IdFreelance    int64
	JobChildCode   string
	JobAddress     string
	JobLong        float64
	JobLat         float64
	JobDescription string
	StartAt        time.Time
	FinishedAt     time.Time
	IdStatus       int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
