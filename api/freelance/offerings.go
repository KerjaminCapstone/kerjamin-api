package freelance

import (
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/labstack/echo/v4"
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
		Where(`public.order.id_status = ?`, 5).
		Scan(&orders)

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
	db.Model(&model.Order{}).Select(`public.order.id_order as id_order_fr, public.order.id_client, public.order.id_freelance, 
				public.order.job_description as keluhan, public.user.no_wa as no_wa_client,
				public.job_child_code.job_child_name as job_title, public.client_data.id_user, 
				public.user.name as client_name, public.order_status.status_name as status`).
		Where(`public.order.id_freelance = ?`, fr.IdFreelance).
		Where(`public.order.id_order = ?`, idOrder).
		Joins(`left join public.client_data on public.client_data.id_client = public.order.id_client`).
		Joins(`left join public.freelance_data on public.freelance_data.id_freelance = public.order.id_freelance`).
		Joins(`left join public.user on public.user.id_user = public.client_data.id_user`).
		Joins(`left join public.job_child_code on public.job_child_code.job_child_code = public.order.job_child_code`).
		Joins(`left join public.order_status on public.order_status.id_status = public.order.id_status`).
		Scan(&order)

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
	if err := db.First(&order, "id_order = ?", idOrder).Error; err != nil {
		return err
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
