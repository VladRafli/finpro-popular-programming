package service

import (
	"my_kelurahan/rt/dto"
	"my_kelurahan/database"
	"my_kelurahan/helpers"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var (
	logger = helpers.InitLogger()
	db     = helpers.ConnectDatabase(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
)

func Create(c echo.Context) (interface{}, *echo.HTTPError) {
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
	body := dto.CreateRtDto{}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind RT.")
	}

	if err := c.Validate(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to validate RT.")
	}

	rt := &database.RT{
		NoRT:  body.NoRT,
		Nama:  body.Nama,
		Alamat: body.Alamat,
	}

	if result := db.Create(rt); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create new RT due internal server error.")
	}

	loggerWithFields.Info("Successfully created new RT.")

	return rt, nil
}

func ReadAll(c echo.Context) (interface{}, *echo.HTTPError) {
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

	rts := &[]database.RT{}

	if result := db.Limit(take).Offset(skip).Find(rts); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to get all RT.")
	}

	loggerWithFields.Info("Successfully get all RT.")

	return rts, nil
}

func Read(c echo.Context) (interface{}, *echo.HTTPError) {
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
	body := dto.ReadRtDto{}
	id := c.Param("id")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "User id is required.")
	}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind RT.")
	}

	rt := &database.RT{
		ID: id,
		NoRT: body.NoRT,
		Nama: body.Nama,
		Alamat: body.Alamat,
	}

	if result := db.Find(rt); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RT not found.")
	}

	loggerWithFields.Info("Successfully get specific RT.")

	return rt, nil
}

func Update(c echo.Context) (interface{}, *echo.HTTPError) {
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
	body := dto.UpdateRtDto{}
	id := c.Param("id")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RT id is required.")
	}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind RT.")
	}

	rt := &database.RT{
		ID: id,
	}

	if result := db.Find(rt); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RT not found.")
	}

	if body.NoRT != "" {
		rt.NoRT = body.NoRT
	}
	if body.Nama != "" {
		rt.Nama = body.Nama
	}
	if body.Alamat != "" {
		rt.Alamat = body.Alamat
	}

	if result := db.Save(rt); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to update RT due to internal server error.")
	}

	loggerWithFields.Info("Successfully update spesific RT.")

	return rt, nil
}

func Delete(c echo.Context) (interface{}, *echo.HTTPError) {
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
		return nil, echo.NewHTTPError(http.StatusBadRequest, "User id is required.")
	}

	rt := &database.RT{
		ID: id,
	}

	if result := db.First(rt); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RT not found.")
	}

	if result := db.Delete(rt); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete RT due to internal server error.")
	}

	loggerWithFields.Info("Successfully delete spesific RT.")

	return rt, nil
}