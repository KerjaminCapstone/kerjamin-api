package schema

type ParentNlpApiResponse struct {
	Data NlpApiResponse `json:"data"`
}

type NlpApiResponse struct {
	NlpScore       float64 `json:"nlp_score"`
	RatingModelSum float64 `json:"rating_model_sum"`
}
