package api

import (
	"log"
	"net/http"
	"picadosYa/internal/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *API) GetFields(c echo.Context) error {
	ctx := c.Request().Context()
	fields, err := a.fieldService.GetFields(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
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
	month, err := time.Parse("2006-01", c.QueryParam("month"))
	if err != nil {
		// Si ocurre un error en la conversión, se responde con un error 400
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid month format"})
	}

	field, err := a.fieldService.GetField(ctx, id, month)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, field)
}

func (a *API) CreateField(c echo.Context) error {
	ctx := c.Request().Context()
	field := new(models.Field)
	if err := c.Bind(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println("Bind data successful")
	if err := a.dataValidator.Struct(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println("Validation successful")
	if err := a.fieldService.SaveField(ctx, field); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Println("Save field successful")
	return c.JSON(http.StatusCreated, field)
}

func (a *API) UpdateField(c echo.Context) error {
	ctx := c.Request().Context()
	field := new(models.Field)
	if err := c.Bind(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := a.dataValidator.Struct(field); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := a.fieldService.UpdateField(ctx, field); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, field)
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
	return c.NoContent(http.StatusNoContent)
}
