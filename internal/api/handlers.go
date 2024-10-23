package api

import (
	"log"
	"net/http"

	"picadosYa/encryption"
	"picadosYa/internal/api/dtos"
	"picadosYa/internal/service"

	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	// valida lo que tenemos asignado en el dto
	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.FirstName, params.Lastname, params.Email, params.Password, params.Phone, params.ProfilePictureUrl, params.Role, params.PositionPlayer)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "user already exists"})
		}
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}
	userCreated := dtos.RegisteredUser{
		FirstName:         params.FirstName,
		LastName:          params.Lastname,
		Email:             params.Email,
		Phone:             params.Phone,
		ProfilePictureUrl: params.ProfilePictureUrl,
		Role:              params.Role,
		PositionPlayer:    params.PositionPlayer,
	}
	return c.JSON(http.StatusCreated, userCreated)
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}
	userCreated := dtos.RegisteredUser{
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		Email:             u.Email,
		Phone:             u.Phone,
		ProfilePictureUrl: u.ProfilePictureUrl,
		Role:              u.Role,
		PositionPlayer:    u.PositionPlayer,
	}
	token, err := encryption.SignedLoginToken(u)
	log.Println(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, userCreated)
}
