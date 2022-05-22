package client

import (
	"net/http"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SubmitOrder(c echo.Context) error {
	form := new(schema.OrderSubmit)
	idFreelance := c.Param("id_freelance")

	if err := c.Bind(form); err != nil {
		return err
	}

	if err := c.Validate(form); err != nil {
		return err
	}

	db := database.GetDBInstance()
	uId, _ := helper.ExtractToken(c)
	user, errUid := helper.FindByUId(uId)
	if errUid != nil {
		return errUid
	}
	timeNow := time.Now()

	idOd := helper.RandomStr(8)
	if idOd == "" {
		return echo.ErrInternalServerError
	}

	clientData, errClient := user.FindClientAcc()
	if errClient != nil {
		return gorm.ErrRecordNotFound
	}

	var freelanceData model.FreelanceData
	if err := db.Where("id_freelance = ?", idFreelance).First(&freelanceData).Error; err != nil {
		return err
	}

	// Nanti disini bakalan ditambahkan api cari address (text)
	// yang didapatkan dari api Google Map
	// Param: longitude latitude, Response: Alamat

	errOrder := db.Create(&model.Order{
		IdOrder:        "OD-" + idOd,
		IdClient:       clientData.IdClient,
		IdFreelance:    freelanceData.IdFreelance,
		JobChildCode:   freelanceData.JobChildCode,
		JobLong:        form.JobLong,
		JobLat:         form.JobLat,
		JobDescription: form.JobDescription,
		AlreadyPaid:    false,
		IdStatus:       1,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}).Error

	if errOrder != nil {
		return errOrder
	}

	res := static.ResponseCreate{
		Message: "Order berhasil dibuat.",
	}
	return c.JSON(http.StatusCreated, res)
}
