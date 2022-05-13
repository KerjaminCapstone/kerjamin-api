package middleware

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/labstack/echo/v4"
)

func CheckRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			uId, role1 := helper.ExtractToken(c)
			var scRole model.Role
			db := database.GetDBInstance()
			// Kalau role yg di jwt sama kaya yang di param CheckRole
			if db.First(&scRole, "id_role = ?", role1); scRole.Name == role {
				// Kalau user nggapunya role yang ada di jwt
				if u, err := helper.IsUserExist(uId); err != nil ||
					!u.HaveRole(role1) {
					return echo.ErrNotFound
				}

				return next(c)
			}

			return &static.Unauthorized{}
		}
	}
}
