package service

import (
	"my_kelurahan/penduduk/dto"
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
	body := dto.CreatePendudukDto{}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind Penduduk.")
	}

	if err := c.Validate(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to validate Penduduk.")
	}

	penduduk := &database.Penduduk{
		NIK:             body.NIK,
		Nama:            body.Nama,
		TempatLahir:     body.TempatLahir,
		TanggalLahir:    body.TanggalLahir,
		Agama:           body.Agama,
		Pekerjaan:       body.Pekerjaan,
		Pendidikan:      body.Pendidikan,
		StatusKawin:     body.StatusKawin,
		StatusHubungan:  body.StatusHubungan,
		Kewarganegaraan: body.Kewarganegaraan,
	}

	if result := db.Create(penduduk); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create new Penduduk due internal server error.")
	}

	loggerWithFields.Info("Successfully created new Penduduk.")

	return penduduk, nil
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

	penduduk := []database.Penduduk{}

	if result := db.Limit(take).Offset(skip).Find(&penduduk); result.Error != nil {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve all Penduduk due internal server error.")
	}

	loggerWithFields.Info("Successfully retrieved all Penduduk.")

	return penduduk, nil
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
	body := dto.ReadPendudukDto{}
	id := c.Param("nik")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "NIK is required.")
	}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind Penduduk.")
	}

	penduduk := database.Penduduk{
		NIK: id,
		Nama: body.Nama,
		TempatLahir: body.TempatLahir,
		TanggalLahir: body.TanggalLahir,
		Agama: body.Agama,
		Pekerjaan: body.Pekerjaan,
		Pendidikan: body.Pendidikan,
		StatusKawin: body.StatusKawin,
		StatusHubungan: body.StatusHubungan,
		Kewarganegaraan: body.Kewarganegaraan,
	}

	if result := db.First(&penduduk); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve Penduduk due internal server error.")
	}

	loggerWithFields.Info("Successfully get specific Penduduk.")

	return penduduk, nil
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
	body := dto.UpdatePendudukDto{}
	id := c.Param("nik")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "NIK is required.")
	}

	if err := c.Bind(&body); err != nil {
		loggerWithFields.Error(err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to bind Penduduk.")
	}

	penduduk := &database.Penduduk{
		NIK: id,
		Nama: body.Nama,
		TempatLahir: body.TempatLahir,
		TanggalLahir: body.TanggalLahir,
		Agama: body.Agama,
		Pekerjaan: body.Pekerjaan,
		Pendidikan: body.Pendidikan,
		StatusKawin: body.StatusKawin,
		StatusHubungan: body.StatusHubungan,
		Kewarganegaraan: body.Kewarganegaraan,
	}

	if result := db.Save(penduduk); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to update Penduduk due internal server error.")
	}

	loggerWithFields.Info("Successfully updated Penduduk.")

	return penduduk, nil
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
	id := c.Param("nik")

	if id == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "NIK is required.")
	}

	penduduk := database.Penduduk{
		NIK: id,
	}

	if result := db.First(&penduduk); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete Penduduk due internal server error.")
	}

	if result := db.Delete(&penduduk); result.Error != nil || result.RowsAffected < 1 {
		loggerWithFields.Error(result.Error)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete Penduduk due internal server error.")
	}

	loggerWithFields.Info("Successfully deleted Penduduk.")

	return penduduk, nil
}