package helper

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
)

func FindByEmail(email string) (*model.User, error) {
	db := database.GetDBInstance()
	var user model.User
	search := db.Where("email = ?", email).First(&user)
	if search.Error != nil {
		return nil, search.Error
	}

	return &user, nil
}
