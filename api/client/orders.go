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

	newIdOd := "OD-" + idOd
	errOrder := db.Create(&model.Order{
		IdOrder:        newIdOd,
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

	res := static.ResponseSuccess{
		Data: struct {
			id_order string
		}{
			id_order: "OD-" + idOd,
		},
	}
	return c.JSON(http.StatusCreated, res)
}

func DetailPesanan(c echo.Context) error {
	idOrder := c.Param("id_order")

	db := database.GetDBInstance()

	var order schema.OrderDetail
	res := db.Model(&model.Order{}).Select(`public.order.job_description, public.freelance_data.rating,
			public.user.name, public.user.no_wa, public.job_child_code.job_child_name, public.order_status.status_name,
			public.order_payment.value_clean, public.order_payment.value_total`).
		Where(`public.order.id_order = ?`, idOrder).
		Joins(`left join public.freelance_data on public.freelance_data.id_freelance = public.order.id_freelance`).
		Joins(`left join public.user on public.user.id_user = public.freelance_data.id_user`).
		Joins(`left join public.job_child_code on public.job_child_code.job_child_code = public.order.job_child_code`).
		Joins(`left join public.order_status on public.order_status.id_status = public.order.id_status`).
		Joins(`left join public.order_payment on public.order_payment.id_order = public.order.id_order`).
		Scan(&order)

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if res.Error != nil {
		return res.Error
	}

	return c.JSON(http.StatusOK, order)
}
