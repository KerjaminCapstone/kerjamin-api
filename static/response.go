package static

type ResponseSuccess struct {
	Data interface{} `json:"data"`
}

type ResponseCreate struct {
	Message string `json:"message"`
}

type ResponseFail struct {
	ErrorMessage string
}
