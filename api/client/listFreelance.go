package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/labstack/echo/v4"
)

type Payload struct {
	JobCode      string  `json:"job_code"`
	JobChildCode string  `json:"job_child_code"`
	AddressLong  float64 `json:"address_long"`
	AddressLat   float64 `json:"address_lat"`
}
type Response struct {
	IdFreelance int `gorm:"primaryKey;autoIncrement;"`
	IdUser      string
	IsTrainee   bool
	Rating      float64
	JobDone     int
	DateJoin    time.Time
	Address     string
	AddressLong float64
	AddressLat  float64
	IsMale      bool
	Dob         time.Time
	Nik         string
	ProfilePict string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Distance    float64 `json:"distance"`
}

// parameter longitude, latitude,
func ListFreelance(c echo.Context) error {

	var payload Payload
	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return echo.ErrBadRequest
	}

	var result []Response

	db := database.GetDBInstance()
	err := db.Raw(`SELECT u."name" , fd.rating , fd.profile_pict , fd.job_done , fd.points , fd.is_trainee,fd.is_male , jcc.job_child_name ,(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) as distance 
	from freelance_data fd, job_child_code jcc ,job_code jc  , "user" u 
	where distance <10 and jcc.job_code  = jc.job_code and fd.job_child_code =jcc.job_child_code and u.id_user = fd.id_user 
	order by distance asc`, payload.AddressLat, payload.AddressLong, payload.AddressLat).Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, result)
}
