package static

type ResponseSuccess struct {
	Data interface{} `json:"data"`
}

type ResponseFail struct {
	Status       int
	Error        bool
	ErrorMessage string
}
