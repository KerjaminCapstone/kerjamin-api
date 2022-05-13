package helper

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
)

func FindRoleByName(name string) (*model.Role, error) {
	db := database.GetDBInstance()
	var role model.Role
	search := db.Where("name = ?", name).First(&role)
	if search.Error != nil {
		return nil, search.Error
	}

	return &role, nil
}
