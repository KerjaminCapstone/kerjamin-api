package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/helper"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
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
	Distance      string    `json:"distance"`
	JobChildName  string    `json:"job_child_name"`
	Address       string    `json:"alamat"`
	Latitude      float64   `json:"address_lat"`
	Longitude     float64   `json:"address_long"`
}

func ListFreelance(c echo.Context) error {

	job_code := c.Param("job_code")
	var result []Response

	type coordinate struct {
		AddressLat  float64 `json:"address_lat"`
		AddressLong float64 `json:"address_long"`
	}
	var clientLatLong coordinate
	userID, _ := helper.ExtractToken(c)
	db := database.GetDBInstance()
	errClient := db.Raw(`select address_lat, address_long from client_data where id_user = ?`, userID).Scan(&clientLatLong).Error
	if errClient != nil {
		return echo.ErrInternalServerError
	}
	// sort_by 1 / 2/ 3 /4
	// 1 == by distance || 2== by rating

	err := db.Raw(`SELECT fd.date_join,fd.job_done, case when fd.is_male = true then 'Pria' else 'Wanita' end as jenis_kelamin,
	fd.id_freelance,u."name" , fd.rating   , fd.points , fd.is_trainee , jcc.job_child_name , fd.address as alamat, fd.address_lat ,fd.address_long
	from freelance_data fd, job_child_code jcc ,job_code jc  , "user" u 
	where jcc.job_code  = jc.job_code and fd.job_child_code =jcc.job_child_code and u.id_user = fd.id_user and jc.job_code=? and
	(6371 * acos( cos( radians(fd.address_lat) ) * cos( radians( ? ) ) *cos( radians( ? ) - radians(fd.address_long) ) 
	+ sin( radians(fd.address_lat) ) * sin( radians( ? ) )) ) <10
	order by distance asc, fd.rating desc`, job_code, clientLatLong.AddressLat, clientLatLong.AddressLong, clientLatLong.AddressLat).Scan(&result).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	url := `https://maps.googleapis.com/maps/api/distancematrix/json?origins=` + fmt.Sprintf("%f", clientLatLong.AddressLat) + `,` + fmt.Sprintf("%f", clientLatLong.AddressLong) + `&destinations=`
	api_key := "&key=" + os.Getenv("API_KEY")

	for i, data := range result {
		if i == 0 {
			url = url + fmt.Sprintf("%f", data.Latitude) + `,` + fmt.Sprintf("%f", data.Longitude)
		} else {
			url = url + `%7C` + fmt.Sprintf("%f", data.Latitude) + `,` + fmt.Sprintf("%f", data.Longitude)
		}
	}
	url += api_key

	// akses gmaps distance api
	var output model.DistanceMatrixResponse
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resMaps, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resMaps.Body.Close()

	body, err := ioutil.ReadAll(resMaps.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonErr := json.Unmarshal(body, &output)
	if jsonErr != nil {
		return err
	}
	//

	var resultJson []Response
	resultJson = make([]Response, 0)
	for _, data := range output.Rows {
		for i, outputData := range data.Elements {
			result[i].Distance = outputData.Distance.HumanReadable
			if outputData.Distance.Meters <= 10000 {
				resultJson = append(resultJson, result[i])
			}
		}
	}

	res := static.ResponseSuccess{
		Error: false,
		Data:  url,
	}
	return c.JSON(http.StatusOK, res)
}
