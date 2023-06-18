package service

import (
	"my_kelurahan/domisili/dto"
	"my_kelurahan/database"
	"my_kelurahan/helpers"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Create(c echo.Context) (interface{}, *echo.HTTPError) {
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
	body := dto.CreateDomisiliDto{}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind Domisili.")
	}

	if err := c.Validate(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to validate Domisili.")
	}

	domisili := &database.Domisili{
		Alamat: body.Alamat,
	}

	if result := db.Create(domisili); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create new Domisili due internal server error.")
	}

	loggerWithFields.Info("Successfully created new Domisili.")

	return domisili, nil
}

func ReadAll(c echo.Context) (interface{}, *echo.HTTPError) {
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
	var (
		take int
		skip int
	)

	takeQuery := c.QueryParam("take")

	if takeQuery != "" {
		if val, err := strconv.ParseInt(takeQuery, 10, 32); err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to parse take.")
		} else {
			take = int(val)
		}
	} else {
		take = -1
	}

	skipQuery := c.QueryParam("skip")

	if skipQuery != "" {
		if val, err := strconv.ParseInt(skipQuery, 10, 32); err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to parse skip.")
		} else {
			skip = int(val)
		}
	} else {
		skip = -1
	}

	domisili := []database.Domisili{}

	if result := db.Limit(take).Offset(skip).Find(&domisili); result.Error != nil {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve all Domisili due internal server error.")
	}

	loggerWithFields.Info("Successfully retrieved all Domisili.")

	return domisili, nil
}

func Read(c echo.Context) (interface{}, *echo.HTTPError) {
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
	body := dto.ReadDomisiliDto{}
	id := c.Param("id")
	
	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Domisili id is required.")
	}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind Domisili.")
	}

	domisili := database.Domisili{
		ID: id,
		Alamat: body.Alamat,
	}

	if result := db.First(&domisili); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusNotFound, "Failed to retrieve Domisili.")
	}

	loggerWithFields.Info("Successfully get specific Domisili.")

	return domisili, nil
}

func Update(c echo.Context) (interface{}, *echo.HTTPError) {
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
	body := dto.UpdateDomisiliDto{}
	id := c.Param("id")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Domisili id is required.")
	}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind Domisili.")
	}

	domisili := &database.Domisili{
		ID: id,
	}

	if result := db.First(domisili); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusNotFound, "Failed to retrieve Domisili.")
	}

	if body.Alamat != "" {
		domisili.Alamat = body.Alamat
	}

	if result := db.Save(domisili); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to update Domisili due internal server error.")
	}

	loggerWithFields.Info("Successfully update spesific Domisili.")

	return domisili, nil
}

func Delete(c echo.Context) (interface{}, *echo.HTTPError) {
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
	id := c.Param("id")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Domisili id is required.")
	}

	domisili := &database.Domisili{
		ID: id,
	}

	if result := db.First(domisili); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusNotFound, "Failed to retrieve Domisili.")
	}

	if result := db.Delete(domisili); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete Domisili due internal server error.")
	}

	loggerWithFields.Info("Successfully delete spesific Domisili.")

	return domisili, nil
}