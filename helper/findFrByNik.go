package helper

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"gorm.io/gorm"
)

func IsFreelancerExist(nik string) (*model.FreelanceData, error) {
	db := database.GetDBInstance()
	var fr model.FreelanceData

	err := db.First(&fr, "nik = ?", nik).Error
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &fr, nil
}
