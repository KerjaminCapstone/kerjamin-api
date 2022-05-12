package auth

import (
	"net/http"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/config"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	UId     string `json:"uid"`
	RoleId1 string `json:"role_id_1"`
	RoleId2 string `json:"role_id_2"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	form := new(schema.Login)

	if err := c.Bind(form); err != nil {
		return err
	}

	if err := c.Validate(form); err != nil {
		return err
	}

	obj, err := helper.FindByEmail(form.Email)

	if err != nil {
		return err
	}
	if helper.CheckPassword(obj.Password, form.Password) {
		return &static.LoginError{}
	}
	roles := obj.FindRoles()
	if len(roles) == 1 {
		roles = append(roles, model.RoleScan{
			IdRole: "",
			IdUser: "",
		})
	}

	claims := &JwtCustomClaims{
		obj.IdUser,
		roles[0].IdRole,
		roles[1].IdUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8760).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, errJwt := token.SignedString(config.GetSignatureKey())

	if errJwt != nil {
		return errJwt
	}

	rsp := static.ResponseToken{
		Token: t,
	}

	return c.JSON(http.StatusOK, rsp)
}
