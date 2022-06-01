package client

import (
	"net/http"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/labstack/echo/v4"
)

type Response struct {
	IdFreelance   int       `json:"id_freelance"`
	Name          string    `json:"name"`
	IsTrainee     bool      `json:"is_trainee"`
	Rating        float64   `json:"rating"`
	JobDone       int       `json:"job_done"`
	DateJoin      time.Time `json:"date_join"`
	Jenis_kelamin string    `json:"jenis_kelamin"`
	// ProfilePict   string    `json:"profile_pict"`
	Distance float64 `json:"distance"`
}

func ListFreelance(c echo.Context) error {

	job_code := c.Param("job_code")
	sort_by := c.Param("sort_by")
	var result []Response

	type coordinate struct {
		AddressLat  float64 `json:"address_lat"`
		AddressLong float64 `json:"address_long"`
	}
	var queryresult coordinate
	userID, _ := helper.ExtractToken(c)
	db := database.GetDBInstance()
	errClient := db.Raw(`select address_lat as lat, address_long as long from client_data where id_user = ?`, userID).Scan(&queryresult).Error
	if errClient != nil {
		return echo.ErrInternalServerError
	}
	// sort_by 1 / 2/ 3 /4
	// 1 == by distance || 2== by rating
	if sort_by == "1" {

		err := db.Raw(`SELECT fd.date_join,fd.job_done, case when fd.is_male = true then 'Pria' else 'Wanita' end as jenis_kelamin, fd.id_freelance,u."name" , fd.rating   , fd.points , fd.is_trainee , jcc.job_child_name ,(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) as distance 
	from freelance_data fd, job_child_code jcc ,job_code jc  , "user" u 
	where jcc.job_code  = jc.job_code and fd.job_child_code =jcc.job_child_code and u.id_user = fd.id_user and jc.job_code=? and
	(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) <10
	order by distance asc, fd.rating desc`, queryresult.AddressLat, queryresult.AddressLong, queryresult.AddressLat, job_code, queryresult.AddressLat, queryresult.AddressLong, queryresult.AddressLat).Scan(&result).Error
		if err != nil {
			return echo.ErrInternalServerError
		}

	} else {

		err := db.Raw(`SELECT fd.date_join,fd.job_done, case when fd.is_male = true then 'Pria' else 'Wanita' end as jenis_kelamin, fd.id_freelance,u."name" , fd.rating   , fd.points , fd.is_trainee , jcc.job_child_name ,(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) as distance 
	from freelance_data fd, job_child_code jcc ,job_code jc  , "user" u 
	where jcc.job_code  = jc.job_code and fd.job_child_code =jcc.job_child_code and u.id_user = fd.id_user and jc.job_code=? and
	(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) <10
	order by rating desc,distance asc`, queryresult.AddressLat, queryresult.AddressLong, queryresult.AddressLat, job_code, queryresult.AddressLat, queryresult.AddressLong, queryresult.AddressLat).Scan(&result).Error
		if err != nil {
			return echo.ErrInternalServerError
		}
	}
	res := static.ResponseSuccess{
		Error: false,
		Data:  result,
	}
	return c.JSON(http.StatusOK, res)
}
