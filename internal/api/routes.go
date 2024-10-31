package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	// ###################
	// User Endpoints
	// ###################
	users := apiGroup.Group("/users")
	users.POST("/register", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	users.GET("/auth/token", a.GetExpiration)
	users.POST("/password-recovery", a.RequestPasswordRecovery) // envía la movida al mail
	users.PUT("/reset-password", a.ResetPassword)

	// ###################
	// Fields Endpoints
	// ###################
	fields := apiGroup.Group("/fields")
	fields.GET("", a.GetFields)
	fields.GET("/:id", a.GetField)
	fields.POST("", a.CreateField)
	fields.PUT("/:id", a.UpdateField)
	fields.DELETE("/:id", a.RemoveField)
}
