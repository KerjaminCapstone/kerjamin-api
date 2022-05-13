package helper

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"gorm.io/gorm"
)

func IsUserExist(uID string) (*model.User, error) {
	db := database.GetDBInstance()
	var user model.User

	err := db.First(&user, "id_user = ?", uID).Error
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}
