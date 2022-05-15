package freelance

import (
	"fmt"
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/labstack/echo/v4"
)

func OfferingList(c echo.Context) error {
	var data []model.Order
	uId, role1 := helper.ExtractToken(c)

	fmt.Println(role1)
	db := database.GetDBInstance()

	err := db.Raw("select * from freelance_data f where f.id_user = ?", uId).Scan(&data).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	result := &static.ResponseSuccess{
		Data: data,
	}

	return c.JSON(http.StatusOK, result)
}
