package client

import (
	"errors"
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DataFreelance(c echo.Context) error {
	idFreelance := c.Param("id_freelance")
	db := database.GetDBInstance()

	var freelanceData model.FreelanceData
	err := db.First(&freelanceData, "id_freelance = ?", idFreelance).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	var user model.User
	err = db.First(&user, "id_user = ?", freelanceData.IdUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	bidang, errBidang := freelanceData.FindFreelanceBidang()
	if errBidang != nil {
		return errBidang
	}
	keahlian, errKeahlian := freelanceData.FindFreelanceKeahlian()
	if errKeahlian != nil {
		return errKeahlian
	}
	nlpTag, errNlp := freelanceData.FindNlpTag()
	if errNlp != nil {
		return errNlp
	}

	data := &schema.FreelanceData{
		Nama:     user.Name,
		Bidang:   bidang,
		Keahlian: keahlian,
		NlpTag:   nlpTag,
	}
	if freelanceData.IsMale {
		data.JenisKelamin = "Pria"
	} else {
		data.JenisKelamin = "Wanita"
	}

	result := &static.ResponseSuccess{
		Data: data,
	}

	return c.JSON(http.StatusOK, result)
}
