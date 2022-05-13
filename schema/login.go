package schema

type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
}
