package api

import (
	"fmt"
	"net/http"
	"picadosYa/encryption"
	"picadosYa/internal/models"
	"strconv"
	"strings"
	"time"

	"bytes"
	"encoding/csv"

	"github.com/jung-kurt/gofpdf"
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

	reservation := new(models.Reservation_without_id)
	if err := c.Bind(reservation); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	id_user, role_user, err := getUserIdAndRole(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseError{Message: err.Error()})
	}
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, reservation.Date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid date format"})
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
		return c.JSON(http.StatusUnauthorized, responseError{Message: "No estás logueado como usuario"})
	}
	if err := a.reservationService.CreateReservation(ctx, &reservationToRegister); err != nil {
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

func (a *API) GetReservationsPerUser(c echo.Context) error {
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

	reservationesFromService, err := a.reservationService.GetReservationsPerUser(ctx, idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, reservationesFromService)
}

func (a *API) GetAllReservationsPerOwner(c echo.Context) error {
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

	reservationesForOwner, err := a.reservationService.GetAllReservationsPerOwner(ctx, idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, reservationesForOwner)
}

func (a *API) GetReservationsPerOwnerExport(c echo.Context) error {
	ctx := c.Request().Context()

	tokenStr := c.Request().Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	claims, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

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

	id_user, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Invalid user ID"})
	}
	idUser := int(id_user)

	reservationsExport, err := a.reservationService.GetAllReservationsExport(ctx, idUser, monthsAgo, hour)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	format := c.QueryParam("format")
	switch format {
	case "pdf":
		return generatePDF(c, reservationsExport)
	case "csv":
		return generateCSV(c, reservationsExport)
	default:
		return c.JSON(http.StatusOK, reservationsExport)
	}
}

func generatePDF(c echo.Context, reservations []models.Reservations_Field_Owner) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Reservations Per Month")
	pdf.Ln(10)

	// Encabezados
	headers := []string{"UserName", "FieldName", "Date", "StartTime", "EndTime", "Type", "Phone", "Status"}
	for _, header := range headers {
		pdf.CellFormat(30, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Filas
	pdf.SetFont("Arial", "", 10)
	for _, reservation := range reservations {
		pdf.CellFormat(30, 10, reservation.User_Name, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, reservation.Field_Name, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, reservation.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, reservation.Start_Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, reservation.End_Time, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, reservation.Type, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, reservation.Phone, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, reservation.Status, "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	// Configurar headers y enviar respuesta
	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().Header().Set("Content-Disposition", "attachment;filename=reservations_per_month.pdf")
	return pdf.Output(c.Response().Writer)
}

func generateCSV(c echo.Context, reservations []models.Reservations_Field_Owner) error {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Encabezados
	writer.Write([]string{"UserName", "FieldName", "Date", "StartTime", "EndTime", "Type", "Phone", "Status"})

	// Filas
	for _, reservation := range reservations {
		writer.Write([]string{
			reservation.User_Name,
			reservation.Field_Name,
			reservation.Date,
			reservation.Start_Time,
			reservation.End_Time,
			reservation.Type,
			reservation.Phone,
			reservation.Status,
		})
	}
	writer.Flush()

	// Configurar headers y enviar respuesta
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", "attachment;filename=reservations_per_month.csv")
	_, err := c.Response().Write(buf.Bytes())
	return err
}

func getUserIdAndRole(c echo.Context) (int, string, error) {
	tokenStr := c.Request().Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	claims, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		return 0, "", err
	}

	id_user, ok1 := claims["id"].(float64)
	role_user, ok2 := claims["role"].(string)
	fmt.Println(claims)
	if !ok1 || !ok2 {
		return 0, "", fmt.Errorf("invalid token claims format")
	}

	idUser := int(id_user)
	return idUser, role_user, nil
}
