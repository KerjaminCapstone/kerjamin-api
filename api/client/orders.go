package client

import (
	"encoding/json"
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
		return echo.ErrNotFound
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
			public.order_payment.value_clean, public.order_payment.value_total, public.order.id_status`).
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

func ConfirmOrder(c echo.Context) error {
	idOrder := c.Param("id_order")

	db := database.GetDBInstance()
	var order model.Order
	res := db.First(&order, "id_order = ?", idOrder)
	if err := res.Error; err != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	timeNow := time.Now()
	order.IdStatus = 6
	order.StartAt = timeNow
	db.Save(&order)

	response := static.ResponseCreate{
		Message: "Order berhasil dikonfirmasi. Order akan segera diproses",
	}
	return c.JSON(http.StatusOK, response)
}

func CancelOrder(c echo.Context) error {
	idOrder := c.Param("id_order")

	db := database.GetDBInstance()
	var order model.Order
	res := db.First(&order, "id_order = ?", idOrder)
	if err := res.Error; err != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	order.IdStatus = 5
	db.Save(&order)

	response := static.ResponseCreate{
		Message: "Order telah dibatalkan oleh Client",
	}
	return c.JSON(http.StatusOK, response)
}

func FinishOrder(c echo.Context) error {
	idOrder := c.Param("id_order")

	db := database.GetDBInstance()
	var order model.Order
	res := db.First(&order, "id_order = ?", idOrder)
	if err := res.Error; err != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	timeNow := time.Now()
	order.IdStatus = 7
	order.FinishedAt = timeNow
	db.Save(&order)

	response := static.ResponseCreate{
		Message: "Order telah diselesaikan oleh Client",
	}
	return c.JSON(http.StatusOK, response)
}

func TasksList(c echo.Context) error {
	idOrder := c.Param("id_order")

	db := database.GetDBInstance()
	var order model.Order
	res := db.First(&order, "id_order = ?", idOrder)
	if err := res.Error; err != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	response := static.ResponseSuccess{
		Data: order.GetTasks(),
	}

	return c.JSON(http.StatusOK, response)
}

func HistoryOrder(c echo.Context) error {
	uId, _ := helper.ExtractToken(c)
	user, err := helper.FindByUId(uId)
	if err != nil {
		return err
	}

	cl, err := user.FindClientAcc()
	if err != nil {
		return err
	}

	db := database.GetDBInstance()
	var orders []schema.OrderItem
	db.Model(&model.Order{}).Select(`public.order.created_at, public.job_child_code.job_child_name,
			public.user.name, public.order_status.status_name`).
		Where(`public.order.id_client = ?`, cl.IdClient).
		Where(`public.order.id_status IN ?`, []int{3, 5, 7}).
		Joins(`left join public.job_child_code on public.job_child_code.job_child_code = public.order.job_child_code`).
		Joins(`left join public.client_data on public.client_data.id_client = public.order.id_client`).
		Joins(`left join public.user on public.user.id_user = public.client_data.id_user`).
		Joins(`left join public.order_status on public.order_status.id_status = public.order.id_status`).
		Scan(&orders)

	if orders == nil {
		orders = []schema.OrderItem{}
	}

	result := &static.ResponseSuccess{
		Data: orders,
	}

	return c.JSON(http.StatusOK, result)
}

func ReviewOrder(c echo.Context) error {
	type review_order struct {
		id_order   string `json:"id_order"`
		rating     int    `json:"rating"`
		commentary string `json:"string"`
	}
	var payload review_order
	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return echo.ErrBadRequest
	}
	db := database.GetDBInstance()
	var id_freelance string

	err := db.Raw(`select id_freelance from "order" where id_order = ?`, payload.id_order).Scan((&id_freelance)).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	err2:= db.Raw (`insert into order_review(id_order,id_freelance,rating,commentary,created_at,updated_at) values(?,?,?,?,?,?)`
	payload.id_order,id_freelance,payload.rating,payload.commentary,time.Now(),time.Now()).Error
	if err2 != nil{
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, result)
}