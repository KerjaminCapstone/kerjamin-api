package freelance

import (
	"fmt"
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
	fmt.Println(fr)
	if err != nil {
		return err
	}

	db := database.GetDBInstance()
	var orders []schema.OfferingItem
	db.Model(&model.Order{}).Select(`public.order.id_client, public.order.id_freelance, public.order.created_at as at, 
				public.job_child_code.job_child_name as job_title, public.client_data.id_user, public.user.name as client_name`).
		Joins(`left join public.client_data on public.client_data.id_client = public.order.id_client`).
		Joins(`left join public.freelance_data on public.freelance_data.id_freelance = public.order.id_freelance`).
		Joins(`left join public.user on public.user.id_user = public.client_data.id_user`).
		Joins(`left join public.job_child_code on public.job_child_code.job_child_code = public.order.job_child_code`).
		Scan(&orders)

	result := &static.ResponseSuccess{
		Data: orders,
	}

	return c.JSON(http.StatusOK, result)
}
