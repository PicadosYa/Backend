package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"picadosYa/internal/models"
)

// GetReservations obtiene una lista de reservas
// @Summary Obtiene una lista de reservas
// @Description Devuelve una lista paginada de reservas
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param limit query int false "Número de reservas a obtener"
// @Param offset query int false "Desplazamiento para paginación"
// @Success 200 {array} models.Reservation
// @Failure 500 {object} responseError
// @Router /reservations [get]
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


// GetReservation obtiene una reserva por ID
// @Summary Obtiene una reserva
// @Description Devuelve una reserva específica por ID
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param id path int true "ID de la reserva"
// @Success 200 {object} models.Reservation
// @Failure 400 {object} responseMessage
// @Failure 500 {object} responseError
// @Router /reservations/{id} [get]
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


// CreateReservation crea una nueva reserva
// @Summary Crea una reserva
// @Description Crea una nueva reserva en el sistema
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param reservation body models.Reservation true "Reserva a crear"
// @Success 201 {object} models.Reservation
// @Failure 400 {object} responseMessage
// @Failure 500 {object} responseError
// @Router /reservations [post]
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


// UpdateReservation actualiza una reserva
// @Summary Actualiza una reserva
// @Description Actualiza una reserva existente
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param reservation body models.Reservation true "Reserva a actualizar"
// @Success 200 {object} models.Reservation
// @Failure 400 {object} responseMessage
// @Failure 500 {object} responseError
// @Router /reservations [put]
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


// DeleteReservation elimina una reserva por ID
// @Summary Elimina una reserva
// @Description Elimina una reserva específica por ID
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param id path int true "ID de la reserva"
// @Success 204
// @Failure 400 {object} responseMessage
// @Failure 500 {object} responseError
// @Router /reservations/{id} [delete]
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
