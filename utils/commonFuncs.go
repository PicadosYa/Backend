package utils

import (
	"fmt"
	"log"
	"picadosYa/internal/models"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
)

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
			keyValue := strings.Split(field, ":")

			key := strings.TrimSpace(keyValue[0])

			value := strings.TrimSpace(keyValue[1])

			switch key {
			case "date":
				parsedDate, _ := time.Parse("2006-01-02", value)
				res.Date = strfmt.Date(parsedDate)
			case "start_time":
				log.Println(value)
				res.StartTime, _ = time.Parse("15:04:00", fmt.Sprintf("%s:00:00", value))
			case "end_time":
				log.Println(value)
				res.EndTime, _ = time.Parse("15:04:00", fmt.Sprintf("%s:00:00", value))
			}
		}
		reservationList = append(reservationList, res)
	}
	return reservationList
}
