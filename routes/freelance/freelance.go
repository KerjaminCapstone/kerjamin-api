package freelance

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/api/freelance"
	"github.com/labstack/echo/v4"
)

func FreelanceSubRoute(group *echo.Group) {
	group.GET("/offerings", freelance.OfferingList)
	group.GET("/offerings/:id_order", freelance.OfferingDetail)
	group.PATCH("/offerings/:id_order/update", freelance.UpdateOffering)
}
