package model

import "time"

type Order struct {
	IdOrder        string `gorm:"primaryKey"`
	IdClient       int
	IdFreelance    int
	JobChildCode   string
	JobAddress     string
	JobLong        float64
	JobLat         float64
	JobDescription string
	StartAt        time.Time
	FinishedAt     time.Time
	AlreadyPaid    bool
	IdStatus       int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
