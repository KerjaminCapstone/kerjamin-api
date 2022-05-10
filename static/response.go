package static

type ResponseSuccess struct {
	Status int
	Error  bool
	Data   string
}

type ResponseFail struct {
	Status       int
	Error        bool
	ErrorMessage string
}
