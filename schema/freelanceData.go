package schema

import "github.com/KerjaminCapstone/kerjamin-backend-v1/model"

type FreelanceData struct {
	Bidang       string               `json:"bidang"`
	Keahlian     string               `json:"keahlian"`
	Nama         string               `json:"nama"`
	Alamat       string               `json:"alamat"`
	JenisKelamin string               `json:"jenis_kelamin"`
	NlpTag       *model.FreelancerNlp `json:"tag_nlp"`
	Distance     string               `json:"jarak"`
}
