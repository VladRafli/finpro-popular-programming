package service

import (
	"my_kelurahan/rw/dto"
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
	body := dto.CreateRwDto{}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind RT.")
	}

	if err := c.Validate(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to validate RT.")
	}

	rw := &database.RW{
		NoRW:  body.NoRW,
		Nama:  body.Nama,
		Alamat: body.Alamat,
	}

	if result := db.Create(rw); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create new RT due internal server error.")
	}

	loggerWithFields.Info("Successfully created new RW.")

	return rw, nil
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

	rws := []database.RW{}

	if result := db.Limit(take).Offset(skip).Find(&rws); result.Error != nil {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve all RW due internal server error.")
	}

	loggerWithFields.Info("Successfully retrieved all RW.")

	return rws, nil
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
	body := dto.ReadRwDto{}
	id := c.Param("id")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "User id is required.")
	}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind RW.")
	}

	rw := &database.RW{
		ID: id,
		NoRW: body.NoRW,
		Nama: body.Nama,
		Alamat: body.Alamat,
	}

	if result := db.Find(rw); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RW not found.")
	}

	loggerWithFields.Info("Successfully retrieved RW.")

	return rw, nil
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
	body := dto.UpdateRwDto{}
	id := c.Param("id")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RW id is required.")
	}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind RW.")
	}

	rw := &database.RW{
		ID: id,
	}

	if result := db.Find(rw); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RW not found.")
	}

	if body.NoRW != "" {
		rw.NoRW = body.NoRW
	}
	if body.Nama != "" {
		rw.Nama = body.Nama
	}
	if body.Alamat != "" {
		rw.Alamat = body.Alamat
	}

	if result := db.Save(rw); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to update RW due internal server error.")
	}

	loggerWithFields.Info("Successfully updated RW.")

	return rw, nil
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
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RW id is required.")
	}

	rw := &database.RW{
		ID: id,
	}

	if result := db.First(rw); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "RW not found.")
	}

	if result := db.Delete(rw); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete RW due internal server error.")
	}

	loggerWithFields.Info("Successfully deleted RW.")

	return rw, nil
}