package api

import (
	"log"
	"net/http"
	"picadosYa/encryption"
	"picadosYa/internal/api/dtos"
	"picadosYa/internal/service"
	"picadosYa/utils"
	"time"

	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}
type responseError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
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

	err = a.serv.RegisterUser(ctx, params.FirstName, params.Lastname, params.Email, params.Password, params.Phone, params.ProfilePictureUrl, params.Role, params.PositionPlayer, params.Age)
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
		Age:               params.Age,
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
		Age:               u.Age,
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
		Path:     "/",
	}

	c.SetCookie(cookie)
	log.Println(cookie)
	return c.JSON(http.StatusOK, userCreated)
}

func (a *API) ResetPassword(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.ResetPassword{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.ResetPassword(ctx, params.Email, params.Token, params.NewPassword)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "user already exists"})
		}
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Password successfully updated"})
}

func (a *API) GetExpiration(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		if err == http.ErrNoCookie {

			return c.JSON(http.StatusUnauthorized, responseMessage{Message: "No hay cookie"})
		}
	}
	if tokenStr == "" {
		tokenStr = cookie.Value
	}

	tkn, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseMessage{Message: "Error al decodificar la cookie"})
	}
	expirationUnix := int64(tkn["expires"].(float64))
	expirationTime := time.Unix(expirationUnix, 0)

	timeRemaining := time.Until(expirationTime)
	if timeRemaining <= 0 {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "El token ha expirado"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Ok"})
}

func (a *API) RequestPasswordRecovery(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RequestPasswordRecovery{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	// Generate a recovery token
	recoveryToken := utils.GenerateRandomDigits(6)

	// Save the token with an expiration time (e.g., 15 minutes)
	err := a.serv.SavePasswordRecoveryToken(ctx, params.Email, recoveryToken, time.Now().Add(15*time.Minute))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Unable to save recovery token"})
	}

	// Send the recovery email
	err = a.serv.SendRecoveryEmail(params.Email, recoveryToken)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Recovery email sent"})
}
