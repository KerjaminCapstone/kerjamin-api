package static

type error interface {
	Error() string
}

type LoginError struct {
}

func (e *LoginError) Error() string {
	return "Email atau password anda salah"
}
