package api

import (
	"log"
	"net/http"
	"picadosYa/encryption"
	"picadosYa/internal/models"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *API) GetFields(c echo.Context) error {
	ctx := c.Request().Context()
	month := c.QueryParam("month")
	var monthParsed time.Time
	if month == "" {
		monthParsed = time.Now()
	} else {
		var err error
		monthParsed, err = time.Parse("2006-01", month)
		if err != nil {
			// Si ocurre un error en la conversión, se responde con un error 400
			return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid month format"})
		}
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}

	offset := c.QueryParam("offset")
	if offset == "" {
		offset = "0"
	}

	limitParsed, err := strconv.Atoi(limit)
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid limit format, would be an integer (ej: 1)"})
	}

	offsetParsed, err := strconv.Atoi(offset)
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid offset format, would be an integer (ej: 10)"})
	}

	fields, err := a.fieldService.GetFields(ctx, monthParsed, limitParsed, offsetParsed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, fields)
}

func (a *API) GetField(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid ID format"})

	}
	month := c.QueryParam("month")
	var monthParsed time.Time
	if month == "" {
		monthParsed = time.Now()
	} else {
		var err error
		monthParsed, err = time.Parse("2006-01", month)
		if err != nil {
			// Si ocurre un error en la conversión, se responde con un error 400
			return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid month format"})
		}
	}

	field, err := a.fieldService.GetField(ctx, id, monthParsed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, field)
}

func (a *API) CreateField(c echo.Context) error {
	ctx := c.Request().Context()
	tokenStr := c.Request().Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	claims, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}
	id_user, ok1 := claims["id"].(float64)
	if ok1 != true {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Check id_user"})
	}
	idUser := int(id_user)
	field := new(models.Field)
	if err := c.Bind(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fieldNew := models.FieldWithID_User{
		Id:              field.Id,
		Name:            field.Name,
		Address:         field.Address,
		Neighborhood:    field.Neighborhood,
		Phone:           field.Phone,
		Latitude:        field.Latitude,
		Longitude:       field.Longitude,
		Type:            field.Type,
		Price:           field.Price,
		Description:     field.Description,
		LogoUrl:         field.LogoUrl,
		AverageRating:   field.AverageRating,
		Services:        field.Services,
		CreationDate:    field.CreationDate,
		Photos:          field.Photos,
		AvailableDays:   field.AvailableDays,
		UnvailableDates: field.UnvailableDates,
		Reservations:    field.Reservations,
		ID_User:         idUser,
	}
	log.Println("Bind data successful")
	if err := a.dataValidator.Struct(fieldNew); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println("Validation successful")
	if err := a.fieldService.SaveField(ctx, &fieldNew); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Println("Save field successful")
	return c.NoContent(http.StatusCreated)
}

func (a *API) UpdateField(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid ID format"})

	}
	field := new(models.Field)
	field.Id = id

	if err := c.Bind(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := a.dataValidator.Struct(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := a.fieldService.UpdateField(ctx, field); err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (a *API) PatchField(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid ID format"})

	}
	field := new(models.Field)

	field.Id = id

	if err := c.Bind(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := a.dataValidator.Struct(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := a.fieldService.PatchField(ctx, field); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}

func (a *API) RemoveField(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid ID format"})

	}
	if err := a.fieldService.RemoveField(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}
