package client

import (
	"net/http"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/labstack/echo/v4"
)

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
	long := c.Param("long")
	lat := c.Param("lat")
	job_code := c.Param("job_code")

	var result []Response

	db := database.GetDBInstance()
	err := db.Raw(`SELECT u."name" , fd.rating , fd.profile_pict , fd.job_done , fd.points , fd.is_trainee,fd.is_male , jcc.job_child_name ,(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) as distance 
	from freelance_data fd, job_child_code jcc ,job_code jc  , "user" u 
	where distance <10 and jcc.job_code  = jc.job_code and fd.job_child_code =jcc.job_child_code and u.id_user = fd.id_user and jc.job_code=?
	order by distance asc`, lat, long, lat, job_code).Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, result)
}
