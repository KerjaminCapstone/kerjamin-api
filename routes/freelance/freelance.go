package freelance

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/api/freelance"
	"github.com/labstack/echo/v4"
)

func FreelanceSubRoute(group *echo.Group) {
	group.GET("/offerings", freelance.OfferingList)
}
