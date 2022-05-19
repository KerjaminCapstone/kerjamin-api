package model

import (
	"time"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
)

type FreelanceData struct {
	IdFreelance int `gorm:"primaryKey;autoIncrement;"`
	IdUser      string
	IsTrainee   bool
	Rating      float64
	JobDone     int
	DateJoin    time.Time
	Address     string
	AddressLong float64
	AddressLat  float64
	IsMale      bool
	Dob         time.Time
	Nik         string
	ProfilePict string
	Points      float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (fd *FreelanceData) FindNlpTag() (*FreelancerNlp, error) {
	var nlpTags FreelancerNlp
	db := database.GetDBInstance()
	if err := db.Select("nlp_tag1, nlp_tag2, nlp_tag3, nlp_tag4, nlp_tag5").
		Find(&nlpTags, "id_freelance = ?", fd.IdFreelance).Error; err != nil {
		return nil, err
	}

	return &nlpTags, nil
}
