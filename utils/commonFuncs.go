package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"picadosYa/encryption"
	"picadosYa/internal/models"
	"strconv"
	"strings"

	//"text/template/parse"
	"time"

	"bytes"
	"encoding/csv"

	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"

	"github.com/jung-kurt/gofpdf"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const APIKEY = "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"

func SliceToString(slice []models.Service) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = fmt.Sprintf("%d", v.ID)
	}
	return strings.Join(strSlice, ",")
}

func SplitString(input string) []string {
	if input == "" {
		return nil
	}
	return strings.Split(input, ",")
}

func ParseServices(servicesStr string) []models.Service {
	if servicesStr == "" {
		return nil
	}
	services := []models.Service{}
	serviceNames := strings.Split(servicesStr, ",")

	for _, name := range serviceNames {
		services = append(services, models.Service{
			Name: name,
		})
	}
	return services
}

func ParseUnavailableDates(dates string) []models.UnvailableDates {
	if dates == "" {
		return nil
	}

	var unavailableDates []models.UnvailableDates
	dateRanges := strings.Split(dates, ",")

	for _, dateRange := range dateRanges {
		// Encuentra la posición del primer espacio que separa las dos fechas
		spaceIndex := strings.LastIndex(dateRange, " ")

		if spaceIndex == -1 {
			fmt.Println("Formato de rango de fechas no válido")
			return unavailableDates
		}

		startDate := dateRange[:spaceIndex-11]
		endDate := dateRange[spaceIndex-10:]

		unavailableDates = append(unavailableDates, models.UnvailableDates{
			FromDate: startDate,
			ToDate:   endDate,
		})

	}
	return unavailableDates
}

// Esto ni yo lo entiendo pero funciona xd
// PD: GPT si lo entiende (supongo)
// PD 2: Es para parsear del raw de sql a un objeto ReservationReduced
func ParseReservations(reservations string) []models.ReservationReduced {
	if reservations == "" {
		return nil
	}
	var reservationList []models.ReservationReduced
	reservationEntries := strings.Split(reservations, "},{")
	for _, entry := range reservationEntries {
		entry = strings.Trim(entry, "{}")
		fields := strings.Split(entry, ",")
		var res models.ReservationReduced
		for _, field := range fields {
			separatorIndex := strings.Index(field, ":")

			key := strings.TrimSpace(field[:separatorIndex])

			value := strings.TrimSpace(field[separatorIndex+1:])

			switch key {
			case "date":
				parsedDate, _ := time.Parse("2006-01-02", value)
				res.Date = strfmt.Date(parsedDate)
			case "start_time":
				log.Println(value)
				parsedTime, _ := time.Parse("15:04:05", value)
				res.StartTime = models.HourMinute(parsedTime)
			case "end_time":
				log.Println(value)
				parsedTime, _ := time.Parse("15:04:05", value)
				res.EndTime = models.HourMinute(parsedTime)

			}
		}
		reservationList = append(reservationList, res)
	}
	return reservationList
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		idx := rand.Int63() % int64(len(letterBytes))
		sb.WriteByte(letterBytes[idx])
	}
	return sb.String()
}

func GenerateRandomDigits(n int) string {
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < n; i++ {
		digit := rand.Intn(10) // genera un número aleatorio entre 0 y 9
		result += strconv.Itoa(digit)
	}
	return result
}

// Sendgrid
func SendEmail(templateID, email, token, name string) error {
	message := mail.NewV3Mail()
	from := mail.NewEmail("picadosya", "picadosya@gmail.com")
	message.SetFrom(from)
	personalization := mail.NewPersonalization()
	to := mail.NewEmail("picadosya", email)
	personalization.AddTos(to)
	personalization.SetDynamicTemplateData("name", name)
	personalization.SetDynamicTemplateData("token", token)
	personalization.SetDynamicTemplateData("email", email)
	message.AddPersonalizations(personalization)
	message.SetTemplateID(templateID)
	client := sendgrid.NewSendClient(APIKEY)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		return fmt.Errorf("failed to send email, status code: %d", response.StatusCode)
	}
	return nil
}

// Cosas de echo
func GenerateUserID(c echo.Context) int {
	tokenStr := c.Request().Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	claims, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		return 0
	}
	id_user, ok1 := claims["id"].(float64)
	if ok1 != true {
		return 0
	}
	idUser := int(id_user)
	return idUser
}

func GetUserIdAndRole(c echo.Context) (int, string, error) {
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

func GeneratePDF(c echo.Context, reservations []models.Reservations_Field_Owner) error {
	pdf := gofpdf.New("L", "mm", "A4", "") // L = Landscape
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	// Título
	pdf.CellFormat(0, 10, "Reservations Export", "", 1, "C", false, 0, "")

	// Encabezados
	headers := []string{"UserName", "FieldName", "Date", "StartTime", "EndTime", "Type", "Phone", "Status"}
	columnWidths := []float64{40, 40, 30, 25, 25, 30, 40, 25} // Ajustar anchos para que se vean todas las columnas

	for i, header := range headers {
		pdf.CellFormat(columnWidths[i], 10, header, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Datos
	pdf.SetFont("Arial", "", 10)
	for _, reservation := range reservations {
		row := []string{
			reservation.User_Name,
			reservation.Field_Name,
			reservation.Date,
			reservation.Start_Time,
			reservation.End_Time,
			reservation.Type,
			reservation.Phone,
			reservation.Status,
		}

		for i, cell := range row {
			pdf.CellFormat(columnWidths[i], 10, cell, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
	}

	// Configurar la respuesta como PDF
	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().Header().Set("Content-Disposition", "attachment;filename=reservations_export.pdf")
	err := pdf.Output(c.Response().Writer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate PDF")
	}

	return nil
}

func GenerateCSV(c echo.Context, reservations []models.Reservations_Field_Owner) error {
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
