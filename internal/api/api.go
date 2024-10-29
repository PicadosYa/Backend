package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"picadosYa/internal/service"
)

type API struct {
	serv               service.Service
	fieldService       service.FieldService
	reservationService service.ReservationService
	dataValidator      *validator.Validate
}

func New(serv service.Service, fieldService service.FieldService, reservationService service.ReservationService) *API {
	return &API{
		serv:               serv,
		fieldService:       fieldService,
		reservationService: reservationService,
		dataValidator:      validator.New(),
	}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	return e.Start(address)
}
