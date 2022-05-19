package freelance

import (
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {
	uId, _ := helper.ExtractToken(c)
	user, err := helper.FindByUId(uId)
	if err != nil {
		return err
	}

	fr, errFr := user.FindFreelanceAcc()
	if errFr != nil {
		return errFr
	}

	nlpTags, errNlpTag := fr.FindNlpTag()
	if errNlpTag != nil {
		return errNlpTag
	}

	data := schema.FreelanceProfile{
		Nama:      user.Name,
		Email:     user.Email,
		IdUserNik: user.IdUser + " / " + fr.Nik,
		NlpTags:   nlpTags,
		Points:    fr.Points,
	}

	res := static.ResponseSuccess{
		Data: data,
	}

	return c.JSON(http.StatusOK, res)
}
