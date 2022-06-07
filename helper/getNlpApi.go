package helper

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/labstack/echo/v4"
)

func GetNlpPoints(commentary string, rating float64) (*schema.ParentNlpApiResponse, error) {
	url := "https://new-prep-dot-kerjadmin.et.r.appspot.com/predict"

	insertData := map[string]interface{}{"komentar": "Kerja bagus, tapi kurang ramah", "rating": 4}
	jsonValue, _ := json.Marshal(insertData)

	req, errForm := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	if errForm != nil {
		return nil, echo.ErrBadRequest
	}

	client := &http.Client{}
	resp, errReq := client.Do(req)

	if resp.StatusCode != 200 || errReq != nil {
		return nil, echo.ErrInternalServerError
	}

	var res schema.ParentNlpApiResponse
	json.NewDecoder(resp.Body).Decode(&res)

	return &res, nil
}
