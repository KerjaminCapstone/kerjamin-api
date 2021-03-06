package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/database"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/routes"
	"github.com/KerjaminCapstone/kerjamin-backend-v1/static"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	database.Init()

	e := routes.Init()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Custom error message
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("Mohon isi %s", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s bukanlah email yang valid", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s harus lebih besar dari %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s harus lebih kurang dari %s", err.Field(), err.Param())
				}

				break
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			report = echo.NewHTTPError(http.StatusNotFound, "Data tidak ditemukan")
		} else if errors.Is(err, &static.AuthError{}) {
			report = echo.NewHTTPError(http.StatusUnauthorized, "User tidak ditemukan")
		} else if errors.Is(err, echo.ErrInternalServerError) {
			report = echo.NewHTTPError(http.StatusUnauthorized, "Terdapat kesalahan")
		} else if errors.Is(err, echo.ErrNotFound) {
			report = echo.NewHTTPError(http.StatusUnauthorized, "Halaman tidak ditemukan")
		}

		c.Logger().Error(report)
		errObj := static.ResponseError{
			Error:   true,
			Message: fmt.Sprintf("%v", report.Message),
		}
		c.JSON(report.Code, errObj)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

// Validation
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
