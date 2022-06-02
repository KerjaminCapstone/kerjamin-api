package auth

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/api/auth"
	"github.com/labstack/echo/v4"
)

func AuthSubRoute(group *echo.Group) {
	group.POST("/sign-up/client", auth.SignUp)
	group.POST("/sign-up/freelancer", auth.SignUpFr)
	group.POST("/login", auth.Login)
	group.GET("/test", auth.TestMapsApi)

}
