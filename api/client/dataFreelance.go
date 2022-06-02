package client

import (
	"errors"
	"net/http"
	"time"

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
		Alamat:   freelanceData.Address,
		Keahlian: keahlian,
		NlpTag:   nlpTag,
	}
	if freelanceData.IsMale {
		data.JenisKelamin = "Pria"
	} else {
		data.JenisKelamin = "Wanita"
	}

	result := &static.ResponseSuccess{
		Error: false,
		Data:  data,
	}

	return c.JSON(http.StatusOK, result)
}

// // data freelancer by gemi
func DataFreelancer(c echo.Context) error {
	idFreelance := c.Param("id_freelance")
	db := database.GetDBInstance()
	type queryResult struct {
		IdFreelance  int       `json:"id_freelance"`
		IdUser       string    `json:"id_user"`
		IsTrainee    bool      `json:"is_trainee"`
		Rating       float64   `json:"rating"`
		JobDone      int       `json:"job_done"`
		DateJoin     time.Time `json:"date_join"`
		Address      string    `json:"address"`
		AddressLong  float64   `json:"address_long"`
		AddressLat   float64   `json:"address_lat"`
		IsMale       bool      `json:"is_male"`
		Dob          time.Time `json:"dob"`
		Nik          string    `json:"nik"`
		ProfilePict  string    `json:"profile_pict"`
		Points       float64   `json:"points"`
		JobChildCode string    `json:"job_child_code"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	var result queryResult
	err := db.Raw(`select * from freelance_data,"user" u where fd.id_user = u.id_user and fd.id_freelance=?`, idFreelance).Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	// d, err := maps.NewClient(maps.WithAPIKey("Insert-API-Key-Here"))
	// if err != nil {
	// 	log.Fatalf("fatal error: %s", err)
	// }
	// r := &maps.DirectionsRequest{
	// 	Origin:      "Sydney",
	// 	Destination: "Perth",
	// }
	// route, _, err := d.Directions(context.Background(), r)
	// if err != nil {
	// 	log.Fatalf("fatal error: %s", err)
	// }
	// fmt.Println(route)
	res := &static.ResponseSuccess{
		Error: false,
		Data:  result,
	}

	return c.JSON(http.StatusOK, res)
}
