package auth

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SignUp(c echo.Context) error {
	form := new(schema.SignUp)
	if err := c.Bind(form); err != nil {
		return err
	}

	if err := c.Validate(form); err != nil {
		return err
	}
	userExist, _ := helper.FindByEmail(form.Email)
	if userExist != nil {
		msg := static.ResponseCreate{
			Error:   true,
			Message: "Email already exist",
		}
		return c.JSON(http.StatusBadRequest, msg)
	}

	db := database.GetDBInstance()

	timeNow := time.Now()
	newUser := &model.User{
		IdUser:    form.Role + "-" + helper.RandomStr(10),
		Name:      form.Nama,
		Email:     form.Email,
		NoWa:      form.NoWa,
		Password:  helper.GeneratePwd(form.Password),
		CreatedAt: timeNow, UpdatedAt: timeNow,
	}
	_ = db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&newUser)
		return nil
	})

	convertJk, _ := strconv.ParseBool(form.JenisKelamin)
	obj := db.Create(&model.ClientData{
		IdUser:    newUser.IdUser,
		Address:   form.Alamat,
		IsMale:    convertJk,
		Nik:       form.Nik,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	})
	clientRole, _ := helper.FindRoleByName("client")
	uR := db.Create(&model.UserRole{
		IdUser: newUser.IdUser,
		IdRole: clientRole.IdRole,
	})
	if obj.Error != nil || uR.Error != nil {
		return obj.Error
	}

	msg := static.ResponseCreate{
		Error:   false,
		Message: "Pengguna berhasil mendaftar",
	}

	return c.JSON(http.StatusCreated, msg)
}
func TestEndpoint(c echo.Context) error {

	var payload schema.SignUp
	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return echo.ErrBadRequest
	}
	user, _ := helper.FindByEmail(payload.Email)

	if user != nil {
		msg := static.ResponseCreate{
			Error:   false,
			Message: "email sudah ada lur",
		}
		return c.JSON(http.StatusOK, msg)
	}

	msg := static.ResponseCreate{
		Error:   false,
		Message: "email belum ada",
	}

	return c.JSON(http.StatusOK, msg)
}
