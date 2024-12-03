package api

import (
	"net/http"
	"picadosYa/internal/models"
	"picadosYa/utils"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-playground/form/v4"
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
			return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid month format"})
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
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid limit format, would be an integer (ej: 1)"})
	}

	offsetParsed, err := strconv.Atoi(offset)
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid offset format, would be an integer (ej: 10)"})
	}

	fields, err := a.fieldService.GetFields(ctx, monthParsed, limitParsed, offsetParsed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, fields)
}

func (a *API) GetField(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid ID format"})

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
			return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid month format"})
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

	multiPartForm, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{
			Message: "Invalid file upload",
			Error:   err.Error(),
		})
	}

	field := &models.Field{}

	// Bindear automáticamente
	decoder := form.NewDecoder()
	decoder.RegisterCustomTypeFunc(func(values []string) (interface{}, error) {
		if len(values) == 0 {
			return nil, nil
		}
		date, _ := time.Parse("2006-01-02", values[0])
		return strfmt.Date(date), nil
	}, strfmt.Date{})

	err = decoder.Decode(&field, multiPartForm.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{
			Message: "Invalid form data",
			Error:   err.Error(),
		})
	}

	// Save the field
	if err := a.fieldService.SaveField(ctx, field, &multiPartForm.File); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: "Failed to save field",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, field)
}

func (a *API) GetFieldsPerOwner(c echo.Context) error {
	ctx := c.Request().Context()
	idUser := utils.GenerateUserID(c)
	fieldsPerOwner, err := a.fieldService.GetFieldsPerOwner(ctx, idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, fieldsPerOwner)
}

func (a *API) UpdateField(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid ID format"})

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
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (a *API) PatchField(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid ID format"})

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
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid ID format"})

	}
	if err := a.fieldService.RemoveField(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}
