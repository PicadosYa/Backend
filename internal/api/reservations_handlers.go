package api

import (
	"net/http"
	"picadosYa/internal/models"
	"picadosYa/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: "Internal server error", Error: err.Error()})
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
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid ID format"})
	}

	reservation, err := a.reservationService.GetReservation(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: "Internal server error", Error: err.Error()})
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
	reservation := new(models.Reservation_without_id)
	if err := c.Bind(reservation); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	id_user, role_user, err := utils.GetUserIdAndRole(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
	}
	userEmail := a.serv.GetUserEmailByID(ctx, id_user)
	canchaName := a.fieldService.GetFieldIndividually(ctx, reservation.FieldID)
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, reservation.Date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid date format"})
	}
	reservationToRegister := models.Reservation{
		FieldID:   reservation.FieldID,
		UserID:    id_user,
		Date:      parsedDate,
		StartTime: reservation.StartTime,
		EndTime:   reservation.EndTime,
		PaymentID: reservation.PaymentID,
	}
	if role_user != "client" {
		return c.JSON(http.StatusUnauthorized, models.ResponseError{Message: "No estás logueado como usuario"})
	}
	if err := a.reservationService.CreateReservation(ctx, &reservationToRegister); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: "Internal server error", Error: err.Error()})
	}
	utils.SendEmail("d-1dc0f0686db042f08a10d5caa4b80612", userEmail.Email, canchaName.Field_Address, reservation.StartTime)
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
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid request body"})
	}

	if err := a.reservationService.UpdateReservation(ctx, reservation); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: "Internal server error", Error: err.Error()})
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
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid ID format"})
	}

	if err := a.reservationService.DeleteReservation(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseError{Message: "Internal server error", Error: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (a *API) GetReservationsPerUser(c echo.Context) error {
	ctx := c.Request().Context()
	idUser := utils.GenerateUserID(c)

	reservationesFromService, err := a.reservationService.GetReservationsPerUser(ctx, idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, reservationesFromService)
}

func (a *API) GetReservationsPerOwnerExport(c echo.Context) error {
	ctx := c.Request().Context()

	monthsAgoParam := c.QueryParam("MonthsAgo")
	hourParam := c.QueryParam("Hour")

	monthsAgo, err := strconv.Atoi(monthsAgoParam)
	if err != nil {
		monthsAgo = 123
	}

	hour, err := strconv.Atoi(hourParam)
	if err != nil {
		hour = 123
	}

	idUser := utils.GenerateUserID(c)

	reservationsExport, err := a.reservationService.GetAllReservationsExport(ctx, idUser, monthsAgo, hour)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	format := c.QueryParam("format")
	switch format {
	case "pdf":
		return utils.GeneratePDF(c, reservationsExport)
	case "csv":
		return utils.GenerateCSV(c, reservationsExport)
	default:
		return c.JSON(http.StatusOK, reservationsExport)
	}
}
