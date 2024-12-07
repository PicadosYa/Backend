package api

import (
	"picadosYa/encryption"
	//	"picadosYa/internal/entity"
	"picadosYa/internal/middlewares"

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
	users.PUT("/update-user-profile", a.UpdateUserProfileInfo, middlewares.JWTMiddleware([]byte(encryption.Key), a.serv))
	users.GET("/check-info", a.GetUserByID)
	users.POST("/add-favourites", a.CreateOrRemoveFavourite)
	users.GET("/favourites-per-user", a.GetFavouritesPerUser)
	users.GET("/refresh-token", a.RefreshToken, middlewares.JWTMiddleware([]byte(encryption.Key), a.serv))

	// ###################
	// Fields Endpoints
	// ###################
	fields := apiGroup.Group("/fields")
	fields.GET("", a.GetFields)
	fields.GET("/:id", a.GetField)
	fields.POST("", a.CreateField, middlewares.JWTMiddleware([]byte(encryption.Key), a.serv))
	fields.PUT("/:id", a.UpdateField)
	fields.PATCH("/:id", a.PatchField)
	fields.DELETE("/:id", a.RemoveField)
	fields.GET("/per-owner", a.GetFieldsPerOwner)

	// ###################
	// Reservations Endpoints
	// ###################
	reservations := apiGroup.Group("/reservations")
	reservations.GET("", a.GetReservations)
	reservations.GET("/:id", a.GetReservation)
	reservations.POST("", a.CreateReservation)
	reservations.PUT("/:id", a.UpdateReservation)
	reservations.DELETE("/:id", a.DeleteReservation)
	reservations.GET("/reservations-per-user", a.GetReservationsPerUser)
	reservations.GET("/reservations-per-owner", a.GetReservationsPerOwnerExport)

	// ###################
	// Payment Endpoints
	// ###################
	apiGroup.POST("/create_preference", a.PaymentPrincipal)
}
