package schema

import "github.com/KerjaminCapstone/kerjamin-backend-v1/model"

type FreelanceProfile struct {
	Nama      string               `json:"nama"`
	Email     string               `json:"email"`
	IdUserNik string               `json:"id_user_nik"`
	NlpTags   *model.FreelancerNlp `json:"nlp_tags"`
	Points    float64              `json:"points"`
}
