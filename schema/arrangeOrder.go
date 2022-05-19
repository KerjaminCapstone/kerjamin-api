package schema

type ArrangeOrder struct {
	Value     int64    `validate:"required" json:"biaya"`
	TaskDescs []string `validate:"required" json:"tasks"`
}
