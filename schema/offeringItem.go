package schema

type OfferingItem struct {
	IdOrderFr  string `json:"id_order"`
	JobTitle   string `json:"job_title"`
	ClientName string `json:"client_name"`
	At         string `json:"at"`
	Status     string `json:"status"`
}