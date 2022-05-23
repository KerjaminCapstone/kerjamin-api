package client

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/api/client"
	"github.com/labstack/echo/v4"
)

func ClientSubRoute(group *echo.Group) {
	group.GET("/me", client.DataPersonal)
	group.POST("/search/freelance", client.ListFreelance)

	group.GET("/freelance/:id_freelance", client.DataFreelance)
	group.POST("/freelance/:id_freelance/order", client.SubmitOrder)

	group.GET("/orders/:id_order", client.DetailPesanan)
	group.PATCH("/orders/:id_order/confirm", client.ConfirmOrder)
	group.PATCH("/orders/:id_order/cancel", client.CancelOrder)
}
