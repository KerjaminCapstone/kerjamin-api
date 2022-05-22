package model

import (
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
)

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

func (o *Order) GetPayment() *OrderPayment {
	db := database.GetDBInstance()

	var payment OrderPayment
	res := db.First(&payment, "id_order = ?", o.IdOrder)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil
	}

	return &payment
}

func (o *Order) GetTasks() []OrderTask {
	db := database.GetDBInstance()

	var tasks []OrderTask
	db.Find(&tasks, "id_order = ?", o.IdOrder)

	return tasks
}
