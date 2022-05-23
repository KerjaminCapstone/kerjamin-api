package schema

import "github.com/KerjaminCapstone/kerjamin-backend-v1/model"

type OrderArrangement struct {
	ValueClean int64             `json:"harga"`
	Tasks      []model.OrderTask `json:"tasks"`
}
