package client

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/api/client"
	"github.com/labstack/echo/v4"
)

func ClientSubRoute(group *echo.Group) {
	group.GET("/freelance/:id_freelance", client.DataFreelance)
	group.GET("/me", client.DataPersonal)
	group.GET("/me", client.SearchFreelance)

}
