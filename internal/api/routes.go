package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {
	users := e.Group("/users")
	users.POST("/register", a.RegisterUser)
	users.POST("/login", a.LoginUser)

	// ###################
	// Fields Endpoints
	// ###################
	fields := e.Group("/fields")
	fields.GET("", a.GetFields)
	fields.GET("/:id", a.GetField)
	fields.POST("", a.CreateField)
	fields.PUT("/:id", a.UpdateField)
	fields.DELETE("/:id", a.RemoveField)
}
