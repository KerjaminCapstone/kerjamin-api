package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/schema"
	"github.com/labstack/echo/v4"
)

func GetNlpPoints(commentary string, rating float64) (*schema.ParentNlpApiResponse, error) {
	url := "https://ml-api-4-dot-kerjamin-capstone.et.r.appspot.com/predict"

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

func GetNlpTag(idFr int) ([]schema.NlpTagResp, error) {
	url := "https://word-cloud-api-dot-kerjamin-capstone.et.r.appspot.com/predict"

	insertData := map[string]interface{}{"id": idFr}
	jsonValue, _ := json.Marshal(insertData)

	req, errForm := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	if errForm != nil {
		return nil, errForm
	}

	client := &http.Client{}
	resp, errReq := client.Do(req)

	if errReq != nil {
		return nil, errReq
	}
	if resp.StatusCode != 200 {
		return nil, errReq
	}

	var res interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	extractBody := res.(map[string]interface{})
	data := extractBody["data"].([]interface{})
	var resStruct []schema.NlpTagResp
	for _, d := range data {
		sifatItem := d.(map[string]interface{})["sifat"].(string)
		valueItem := d.(map[string]interface{})["value"].(float64)
		item := schema.NlpTagResp{
			Sifat: sifatItem,
			Value: valueItem,
		}
		resStruct = append(resStruct, item)
	}

	sort.Slice(resStruct, func(i, j int) bool {
		return resStruct[i].Value > resStruct[j].Value
	})

	lenMin := 5 - len(resStruct)
	for i := 0; i < lenMin; i++ {
		emptyItem := schema.NlpTagResp{
			Sifat: "",
			Value: 0,
		}
		resStruct = append(resStruct, emptyItem)
	}

	return resStruct, nil
}
