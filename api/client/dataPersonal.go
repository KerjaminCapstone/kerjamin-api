package client

import (
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/labstack/echo/v4"
)

func DataPersonal(c echo.Context) error {
	var result model.ClientData
	uId, _ := helper.ExtractToken(c)

	db := database.GetDBInstance()
	err := db.Raw("select * from client_data where id_user=?", uId).Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, result)
}
