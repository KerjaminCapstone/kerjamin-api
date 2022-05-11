package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DataFreelance(c echo.Context) error {
	idUser := c.QueryParam("idUser")
	db := database.GetDBInstance()

	var user model.User
	err := db.First(&user, "id_user = ?", idUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}
