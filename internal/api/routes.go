package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	// ###################
	// User Endpoints
	// ###################
	users := apiGroup.Group("/users")
	users.POST("/register", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	users.GET("/auth/token", a.GetExpiration)
	users.POST("/password-recovery", a.RequestPasswordRecovery) // envía el correo
	users.PUT("/reset-password", a.ResetPassword)
	users.GET("/verify", a.UpdateVerifyUser)
	users.POST("/verify-user-email", a.VerifyUserEmail) //envía el correo
	users.PUT("/update-user-profile", a.UpdateUserProfileInfo)

	// ###################
	// Fields Endpoints
	// ###################
	fields := apiGroup.Group("/fields")
	fields.GET("", a.GetFields)
	fields.GET("/:id", a.GetField)
	fields.POST("", a.CreateField)
	fields.PUT("/:id", a.UpdateField)
	fields.PATCH("/:id", a.PatchField)
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
	reservations.GET("/reservations-per-user/:id", a.GetReservationsPerUser)
}
