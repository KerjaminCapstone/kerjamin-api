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
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		HousePict   string    `json:"house_pict"`
		NoWa        string    `json:"no_wa"`
		Address     string    `json:"address"`
		AddressLong float64   `json:"address_long"`
		AddressLat  float64   `json:"address_lat"`
		IsMale      bool      `json:"is_male"`
		Dob         time.Time `json:"dob"`
		Nik         string    `json:"nik"`
		ProfilePict string    `json:"profile_pict"`
	}
	var result Result

	uId, _ := helper.ExtractToken(c)

	db := database.GetDBInstance()
	err := db.Raw(`select u.email,u.no_wa,u."name" ,cd.house_pict ,cd.address ,cd.address_long,cd.address_lat ,cd.is_male ,cd.dob ,cd.nik ,cd.profile_pict  
	from "user" u, client_data cd 
	where u.id_user = cd.id_user and u.id_user=?`, uId).Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, result)
}
