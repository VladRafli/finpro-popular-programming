package service

import (
	"my_kelurahan/auth/dto"
	"my_kelurahan/database"
	"my_kelurahan/helpers"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) (interface{}, *echo.HTTPError) {
	logger := helpers.InitLogger()
	db     := helpers.ConnectDatabase(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
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
	body := dto.ReadLoginDto{}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind user.")
	}

	if err := c.Validate(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to validate user.")
	}

	user := &database.Users{
		Email:    body.Email,
		Password: body.Password,
	}

	if result := db.First(user); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "User not found.")
	}

	if error := bcrypt.CompareHashAndPassword([]byte(body.Password), []byte(user.Password)); error != nil {
		loggerWithFields.Error(error)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Wrong credentials.")
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Subject:   user.ID,
		Issuer:    "my_kelurahan",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := tokenClaims.SignedString(os.Getenv("JWT_SECRET"))

	if err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token.")
	}

	loggerWithFields.Info("Successfully logged in.")

	return echo.Map{
		"token": token,
		"user":  user,
	}, nil
}

func Register(c echo.Context) (interface{}, *echo.HTTPError) {
	logger := helpers.InitLogger()
	db     := helpers.ConnectDatabase(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
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
	body := dto.CreateUserDto{}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind user.")
	}

	if err := c.Validate(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to validate user.")
	}

	user := &database.Users{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
		Role:     body.Role,
	}

	if result := db.Create(&user); result.Error != nil {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user due internal server error.")
	}

	loggerWithFields.Info("Successfully registered new user.")

	return user, nil
}

func Logout(c echo.Context) (interface{}, *echo.HTTPError) {
	logger := helpers.InitLogger()
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

	loggerWithFields.Info("Successfully logged out.")

	return nil, nil
}
