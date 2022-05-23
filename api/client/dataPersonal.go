package client

import (
	"net/http"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/labstack/echo/v4"
)

func DataPersonal(c echo.Context) error {
	type Result struct {
		IdClient    int `gorm:"primaryKey;autoIncrement;"`
		IdUser      string
		Name        string
		HousePict   string
		PhoneNumber string
		Address     string
		AddressLong float64
		AddressLat  float64
		IsMale      bool
		Dob         time.Time
		Nik         string
		ProfilePict string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	var result Result

	uId, _ := helper.ExtractToken(c)

	db := database.GetDBInstance()
	err := db.Raw(`select u."name" ,cd.house_pict ,cd.address ,cd.address_long,cd.address_lat ,cd.is_male ,cd.dob ,cd.nik ,cd.profile_pict  
	from "user" u, client_data cd 
	where u.id_user = cd.id_user and u.id_user=?`, uId).Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, result)
}
