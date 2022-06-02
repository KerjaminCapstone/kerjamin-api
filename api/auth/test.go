package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/model"
	"github.com/labstack/echo/v4"
)

func TestMapsApi(c echo.Context) error {

	url := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=Washington,%20DC&destinations=New%20York%20City,%20NY&units=imperial&key=AIzaSyAzDa6dzrwku9v_Dfq0YxwQgZVpCSykG7c"

	var output model.DistanceMatrixResponse

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return echo.ErrInternalServerError
	}

	jsonErr := json.Unmarshal(body, &output)
	if jsonErr != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, output)
}
