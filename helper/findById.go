package helper

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
)

func FindByUId(uId string) (*model.User, error) {
	db := database.GetDBInstance()
	var user model.User
	search := db.Where("id_user = ?", uId).First(&user)
	if search.Error != nil {
		return nil, search.Error
	}

	return &user, nil
}
