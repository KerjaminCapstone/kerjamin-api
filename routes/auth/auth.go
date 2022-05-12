package auth

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/api/auth"
	"github.com/labstack/echo/v4"
)

func AuthSubRoute(group *echo.Group) {
	group.POST("/sign-up", auth.SignUp)
	group.POST("/login", auth.Login)
}
