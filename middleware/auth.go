package middleware

import (
	"github.com/KerjaminCapstone/kerjamin-backend-v1/config"
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: config.GetSignatureKey(),
})
