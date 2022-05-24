package client

import (
	"net/http"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/labstack/echo/v4"
)

type Response struct {
	IdFreelance int       `json:"id_freelance"`
	Name        string    `json:"name"`
	IsTrainee   bool      `json:"is_trainee"`
	Rating      float64   `json:"rating"`
	JobDone     int       `json:"job_done"`
	DateJoin    time.Time `json:"date_join"`
	IsMale      bool      `json:"is_male"`
	Dob         time.Time `json:"dob"`
	ProfilePict string    `json:"profile_pict"`
	Distance    float64   `json:"distance"`
}

// parameter longitude, latitude,
func ListFreelance(c echo.Context) error {
	long := c.Param("long")
	lat := c.Param("lat")
	job_code := c.Param("job_code")

	var result []Response

	db := database.GetDBInstance()
	err := db.Raw(`SELECT fd.date_join, fd.id_freelance,u."name" , fd.rating , fd.profile_pict , fd.job_done , fd.points , fd.is_trainee,fd.is_male , jcc.job_child_name ,(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) as distance 
	from freelance_data fd, job_child_code jcc ,job_code jc  , "user" u 
	where jcc.job_code  = jc.job_code and fd.job_child_code =jcc.job_child_code and u.id_user = fd.id_user and jc.job_code=? and
	(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) <10
	order by distance asc`, lat, long, lat, job_code, lat, long, lat).Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, result)
}
