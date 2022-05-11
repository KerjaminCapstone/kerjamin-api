package auth

import (
	"net/http"
	"strconv"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
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

	db := database.GetDBInstance()
	timeNow := time.Now()
	newUser := &model.User{
		IdUser:    form.Role + "-" + helper.RandomStr(10),
		Username:  form.Username,
		Name:      form.Nama,
		Email:     form.Email,
		Password:  helper.GeneratePwd(form.Password),
		CreatedAt: timeNow, UpdatedAt: timeNow,
	}
	_ = db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&newUser)
		return nil
	})

	convertJk, _ := strconv.ParseBool(form.JenisKelamin)
	convertDate, _ := time.Parse("2006-01-02", form.TanggalLahir)
	if form.Role == "FR" {
		db.Create(&model.FreelanceData{
			IdUser:    newUser.IdUser,
			IsTrainee: false,
			Rating:    0.0,
			JobDone:   0,
			DateJoin:  timeNow,
			Address:   form.Alamat,
			IsMale:    convertJk,
			Dob:       convertDate,
			Nik:       form.Nik,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		})
	} else if form.Role == "CL" {
		db.Create(&model.ClientData{
			IdUser:    newUser.IdUser,
			Address:   form.Alamat,
			IsMale:    convertJk,
			Dob:       convertDate,
			Nik:       form.Nik,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		})
	}

	return c.JSON(http.StatusCreated, newUser.IdUser)
}
