package client

import (
	"encoding/json"
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/labstack/echo/v4"
)

func PaymentMethod(c echo.Context) error {

	type Result struct {
		Id_method    int    `json:"id_method"`
		Payment_name string `json:"payment_name"`
	}
	var result Result
	db := database.GetDBInstance()
	err := db.Raw("select * from payment_method").Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, result)
}

func OrderPayment(c echo.Context) error {
	//post
	type Payload struct {
		Id_order string `json:"id_order"`
	}
	var payload Payload
	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return echo.ErrBadRequest
	}
	db := database.GetDBInstance()
	err := db.Raw("update order_payment set is_paid=true where id_order=?", payload).Scan(&payload).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, payload)
}
