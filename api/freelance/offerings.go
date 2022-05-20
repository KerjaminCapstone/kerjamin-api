package freelance

import (
	"net/http"
	"strconv"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func OfferingList(c echo.Context) error {
	uId, _ := helper.ExtractToken(c)
	user, err := helper.FindByUId(uId)
	if err != nil {
		return err
	}

	fr, err := user.FindFreelanceAcc()
	if err != nil {
		return err
	}

	db := database.GetDBInstance()

	var orders []schema.OfferingItem
	db.Model(&model.Order{}).Select(`public.order.id_order as id_order_fr, public.order.id_client, public.order.id_freelance, public.order.created_at as at, 
				public.job_child_code.job_child_name as job_title, public.client_data.id_user, public.user.name as client_name`).
		Where(`public.order.id_freelance = ?`, fr.IdFreelance).
		Joins(`left join public.client_data on public.client_data.id_client = public.order.id_client`).
		Joins(`left join public.freelance_data on public.freelance_data.id_freelance = public.order.id_freelance`).
		Joins(`left join public.user on public.user.id_user = public.client_data.id_user`).
		Joins(`left join public.job_child_code on public.job_child_code.job_child_code = public.order.job_child_code`).
		Where(`public.order.id_status IN ?`, []int{1, 2, 5}). // Diterima, proses, assigned
		Scan(&orders)
	if orders == nil {
		orders = []schema.OfferingItem{}
	}

	result := &static.ResponseSuccess{
		Data: orders,
	}

	return c.JSON(http.StatusOK, result)
}

func OfferingDetail(c echo.Context) error {
	idOrder := c.Param("id_order")

	uId, _ := helper.ExtractToken(c)
	user, err := helper.FindByUId(uId)
	if err != nil {
		return err
	}

	fr, err := user.FindFreelanceAcc()
	if err != nil {
		return err
	}

	db := database.GetDBInstance()

	var order schema.OfferingDetail
	res := db.Model(&model.Order{}).Select(`public.order.id_order as id_order_fr, public.order.id_client, public.order.id_freelance, 
				public.order.job_description as keluhan, public.user.no_wa as no_wa_client, public.order.job_long, public.order.job_lat,
				public.job_child_code.job_child_name as job_title, public.client_data.id_user, 
				public.user.name as client_name, public.order_status.status_name as status,
				public.order_payment.value as biaya, public.order_review.commentary as komentar,
				public.order_review.rating as rating`).
		Where(`public.order.id_freelance = ?`, fr.IdFreelance).
		Where(`public.order.id_order = ?`, idOrder).
		Joins(`left join public.client_data on public.client_data.id_client = public.order.id_client`).
		Joins(`left join public.freelance_data on public.freelance_data.id_freelance = public.order.id_freelance`).
		Joins(`left join public.user on public.user.id_user = public.client_data.id_user`).
		Joins(`left join public.job_child_code on public.job_child_code.job_child_code = public.order.job_child_code`).
		Joins(`left join public.order_status on public.order_status.id_status = public.order.id_status`).
		Joins(`left join public.order_payment on public.order_payment.id_order = public.order.id_order`).
		Joins(`left join public.order_review on public.order_review.id_order = public.order.id_order`).
		Scan(&order)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	biayaInt, _ := strconv.Atoi(order.Biaya)
	order.Biaya = humanize.Comma(int64(biayaInt))
	result := &static.ResponseSuccess{
		Data: order,
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateOffering(c echo.Context) error {
	idOrder := c.Param("id_order")
	form := new(schema.UpdateOffering)

	if err := c.Bind(form); err != nil {
		return err
	}

	if err := c.Validate(form); err != nil {
		return err
	}

	db := database.GetDBInstance()
	var order model.Order
	res := db.First(&order, "id_order = ?", idOrder)
	if err := res.Error; err != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	order.IdStatus = int64(form.Status)
	if err := db.Save(&order).Error; err != nil {
		return err
	}

	response := &static.ResponseCreate{
		Message: "Status order berhasil diperbarui",
	}

	return c.JSON(http.StatusOK, response)
}

func ArrangeOffering(c echo.Context) error {
	idOrder := c.Param("id_order")
	form := new(schema.ArrangeOrder)

	if err := c.Bind(form); err != nil {
		return err
	}

	if err := c.Validate(form); err != nil {
		return err
	}

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
	newOdPayment := model.OrderPayment{
		IdOrder:   order.IdOrder,
		Value:     int(form.Value),
		IdMethod:  1,
		IsPaid:    false,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	var tasks []model.OrderTask
	for _, t := range form.TaskDescs {
		task := model.OrderTask{
			IdOrder:    order.IdOrder,
			TaskDesc:   t,
			TaskStatus: false,
			CreatedAt:  timeNow,
			UpdatedAt:  timeNow,
		}
		tasks = append(tasks, task)
	}

	err1 := db.Create(&newOdPayment).Error
	if err1 != nil {
		return err1
	}
	err2 := db.Create(&tasks).Error
	if err2 != nil {
		return err2
	}

	response := static.ResponseCreate{
		Message: "Biaya dan pekerjaan berhasil ditentukan",
	}

	return c.JSON(http.StatusCreated, response)
}

func RefreshStatus(c echo.Context) error {
	idOrder := c.Param("id_order")

	db := database.GetDBInstance()

	var status schema.RefreshStatus
	res := db.Model(&model.Order{}).Select(`public.order.id_status, public.order_status.id_status, public.order_status.status_name`).
		Where(`public.order.id_order = ?`, idOrder).
		Joins(`left join public.order_status on public.order_status.id_status = public.order.id_status`).
		Scan(&status)

	if err := res.Error; err != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	response := static.ResponseSuccess{
		Data: status,
	}
	return c.JSON(http.StatusOK, response)
}

func HistoriOffering(c echo.Context) error {
	uId, _ := helper.ExtractToken(c)
	user, err := helper.FindByUId(uId)
	if err != nil {
		return err
	}

	fr, err := user.FindFreelanceAcc()
	if err != nil {
		return err
	}

	db := database.GetDBInstance()

	var orders []schema.OfferingItem
	db.Model(&model.Order{}).Select(`public.order.id_order as id_order_fr, public.order.id_client, public.order.id_freelance, public.order.created_at as at, 
				public.job_child_code.job_child_name as job_title, public.client_data.id_user, public.user.name as client_name`).
		Where(`public.order.id_freelance = ?`, fr.IdFreelance).
		Joins(`left join public.client_data on public.client_data.id_client = public.order.id_client`).
		Joins(`left join public.freelance_data on public.freelance_data.id_freelance = public.order.id_freelance`).
		Joins(`left join public.user on public.user.id_user = public.client_data.id_user`).
		Joins(`left join public.job_child_code on public.job_child_code.job_child_code = public.order.job_child_code`).
		Where(`public.order.id_status IN ?`, []int{3, 4}). // Selesai dan ditolak
		Scan(&orders)
	if orders == nil {
		orders = []schema.OfferingItem{}
	}

	result := &static.ResponseSuccess{
		Data: orders,
	}

	return c.JSON(http.StatusOK, result)
}
