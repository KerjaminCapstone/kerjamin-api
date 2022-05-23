package schema

type OrderSubmit struct {
	JobLong        float64 `json:"job_long" validate:"required"`
	JobLat         float64 `json:"job_lat" validate:"required"`
	JobDescription string  `json:"job_description" validate:"required"`
}
