package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"picadosYa/internal/models"
)

func (a *API) GetReservations(c echo.Context) error {
	ctx := c.Request().Context()
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	reservations, err := a.reservationService.GetReservations(ctx, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, reservations)
}

func (a *API) GetReservation(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid ID format"})
	}

	reservation, err := a.reservationService.GetReservation(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, reservation)
}

func (a *API) CreateReservation(c echo.Context) error {
	ctx := c.Request().Context()
	reservation := new(models.Reservation)
	if err := c.Bind(reservation); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request body"})
	}

	if err := a.reservationService.CreateReservation(ctx, reservation); err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.JSON(http.StatusCreated, reservation)
}

func (a *API) UpdateReservation(c echo.Context) error {
	ctx := c.Request().Context()
	reservation := new(models.Reservation)
	if err := c.Bind(reservation); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request body"})
	}

	if err := a.reservationService.UpdateReservation(ctx, reservation); err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.JSON(http.StatusOK, reservation)
}

func (a *API) DeleteReservation(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid ID format"})
	}

	if err := a.reservationService.DeleteReservation(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
