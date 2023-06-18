package middlewares

import (
	"my_kelurahan/helpers"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var (
	logger = helpers.InitLogger()
)

func ValidateJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			loggerWithFields := logger.WithFields(logrus.Fields{
				"method":        c.Request().Method,
				"url":           c.Request().RequestURI,
				"contentLength": c.Response().Size,
				"status":        c.Response().Status,
				"host":          c.Request().Host,
				"remoteAddr":    c.RealIP(),
				"userAgent":     c.Request().UserAgent(),
				"responseTime":  c.Response().Header().Get("Time"),
			})
		
			token := c.Request().Header.Get("Authorization")
		
			if token == "" {
				loggerWithFields.Error("Token is required.")
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"statusCode": http.StatusUnauthorized,
					"message":    "Token is required.",
				})
			}
		
			result, err := jwt.ParseWithClaims(
				token,
				jwt.StandardClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("JWT_SECRET")), nil
				},
			)
			if err != nil {
				loggerWithFields.Error("Failed to parse token.")
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"statusCode": http.StatusUnauthorized,
					"message":    "Failed to parse token.",
				})
			}
		
			claims, ok := result.Claims.(jwt.StandardClaims)
		
			if !ok {
				loggerWithFields.Error("Failed to parse token.")
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"statusCode": http.StatusUnauthorized,
					"message":    "Failed to parse token.",
				})
			}
		
			if claims.Issuer != "my_kelurahan" {
				loggerWithFields.Error("Invalid token issuer.")
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"statusCode": http.StatusUnauthorized,
					"message":    "Invalid token issuer.",
				})
			}
		
			if claims.ExpiresAt < time.Now().Local().Unix() {
				loggerWithFields.Error("Token is expired.")
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"statusCode": http.StatusUnauthorized,
					"message":    "Token is expired.",
				})
			}

			return next(c)
		}
	}
}
