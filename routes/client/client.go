package client

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/api/client"
	"github.com/labstack/echo/v4"
)

func ClientSubRoute(group *echo.Group) {
	group.GET("/freelance", client.DataFreelance)
	group.GET("/me", client.DataPersonal)
	group.POST("/search/freelance", client.ListFreelance)

}
