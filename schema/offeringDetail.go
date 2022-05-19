package schema

type OfferingDetail struct {
	IdOrderFr  string `json:"id_order"`
	JobTitle   string `json:"job_title"`
	ClientName string `json:"client_name"`
	Keluhan    string `json:"keluhan"`
	NoWaClient string `json:"no_wa_client"`
	Status     string `json:"status"`
	Biaya      string `json:"biaya"`
	Komentar   string `json:"komentar"`
	Rating     string `json:"rating"`
}
