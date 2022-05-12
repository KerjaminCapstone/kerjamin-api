package model

import (
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
)

type User struct {
	IdUser    string `gorm:"primaryKey"`
	Name      string `validate:"required"`
	Username  string `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required,gte=5,lte=30"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleScan struct {
	IdRole string
	IdUser string
}

func (u *User) FindRoles() []RoleScan {
	var roles []RoleScan
	db := database.GetDBInstance()

	db.Model(&User{}).Select("users.id_user, roles.id_role").Where("users.id_user = ?", u.IdUser).
		Joins("left join user_role on user_role.id_user = users.id_user").
		Joins("left join roles on roles.id_role = user_role.id_role").
		Scan(&roles)

	return roles
}
