package freelance

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/api/freelance"
	"github.com/labstack/echo/v4"
)

func FreelanceSubRoute(group *echo.Group) {
	group.GET("/offerings", freelance.OfferingList)
	group.GET("/offerings/:id_order", freelance.OfferingDetail)
	group.PATCH("/offerings/:id_order/update", freelance.UpdateOffering)

	group.POST("/offerings/:id_order/arrangement/task", freelance.AddTask)
	group.DELETE("/offerings/:id_order/arrangement/task/:id_task", freelance.DeleteTask)
	group.POST("/offerings/:id_order/arrangement", freelance.ArrangeOffering)
	group.GET("/offerings/:id_order/arrangement", freelance.GetArrangement)

	group.GET("/offerings/:id_order/status", freelance.RefreshStatus)

	group.GET("/history", freelance.HistoriOffering)

	group.GET("/me", freelance.GetProfile)
}
