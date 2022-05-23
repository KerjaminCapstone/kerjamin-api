package client

import (
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/labstack/echo/v4"
)

// HAVEN'T WORK
func OngoingOrder(c echo.Context) error {
	var result []model.Order
	uId, _ := helper.ExtractToken(c)
	var id_client string
	db := database.GetDBInstance()
	err := db.Raw("select * from client_data where id_user=?", uId).Scan(&id_client).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	err = db.Raw(`select * from "order" o 
	where o.id_client =?`, id_client).Scan(&id_client).Error
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, result)
}
