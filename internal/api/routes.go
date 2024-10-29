package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	users := apiGroup.Group("/users")
	users.POST("/register", a.RegisterUser)
	users.POST("/login", a.LoginUser)

	// ###################
	// Fields Endpoints
	// ###################
	fields := apiGroup.Group("/fields")
	fields.GET("", a.GetFields)
	fields.GET("/:id", a.GetField)
	fields.POST("", a.CreateField)
	fields.PUT("/:id", a.UpdateField)
	fields.DELETE("/:id", a.RemoveField)

	// ###################
	// Reservations Endpoints
	// ###################
	reservations := apiGroup.Group("/reservations")
	reservations.GET("", a.GetReservations)
	reservations.GET("/:id", a.GetReservation)
	reservations.POST("", a.CreateReservation)
	reservations.PUT("/:id", a.UpdateReservation)
	reservations.DELETE("/:id", a.DeleteReservation)
}
