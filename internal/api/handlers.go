package api

import (
	"log"
	"net/http"
	"picadosYa/encryption"
	"picadosYa/internal/api/dtos"
	"picadosYa/internal/service"
	"picadosYa/utils"
	"strings"
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

	err = a.serv.RegisterUser(ctx, params.FirstName, params.Lastname, params.Email, params.Password, params.Phone, params.Role, params.AcceptedTerms)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "user already exists"})
		}
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}
	userCreated := dtos.RegisteredUser{
		FirstName: params.FirstName,
		LastName:  params.Lastname,
		Email:     params.Email,
		Phone:     params.Phone,
		Role:      params.Role,
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
	log.Println(err)
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}
	userCreated := dtos.LoguedUser{
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
		Phone:      u.Phone,
		Role:       u.Role,
		IsVerified: u.IsVerified,
	}
	token, err := encryption.SignedLoginToken(u)
	log.Println(token)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  userCreated,
		"token": token,
	})
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
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	tkn, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		log.Println(tokenStr)
		log.Println(tkn)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	expiresVal, ok := tkn["exp"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Formato de token inválido"})
	}

	expirationUnix := int64(expiresVal)
	expirationTime := time.Unix(expirationUnix, 0)

	timeRemaining := time.Until(expirationTime)
	if timeRemaining <= 0 {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "El token ha expirado"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Ok"})
}

func (a *API) VerifyUserEmail(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RequestSendEmail{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	// Genera el token de 6 digitos
	recoveryToken := utils.GenerateRandomDigits(6)

	err := a.serv.SaveToken(ctx, params.Email, recoveryToken, time.Now().Add(15*time.Minute))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Unable to save recovery token"})
	}

	// Envía el mail
	err = a.serv.SendVerifyEmail(params.Email, recoveryToken)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Verify email sent"})
}

func (a *API) UpdateVerifyUser(c echo.Context) error {
	ctx := c.Request().Context()
	token := c.QueryParam("token")

	// Verificar el token en la base de datos
	user, err := a.serv.GetUserByToken(ctx, token)
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	// Actualizar isVerified a true
	err = a.serv.UpdateUserVerification(ctx, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "User updated successfully"})
}

func (a *API) RequestPasswordRecovery(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RequestSendEmail{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	// Genera el token de 6 digitos
	recoveryToken := utils.GenerateRandomDigits(6)

	err := a.serv.SaveToken(ctx, params.Email, recoveryToken, time.Now().Add(15*time.Minute))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Unable to save recovery token"})
	}

	// Envía el mail
	err = a.serv.SendRecoveryEmail(params.Email, recoveryToken)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "Recovery email sent"})
}

func (a *API) UpdateUserProfileInfo(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.UpdateUser{}
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	// valida lo que tenemos asignado en el dto
	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.UpdateUserInfo(ctx, params.FirstName, params.LastName, params.Email, params.Phone, params.PositionPlayer, params.TeamName, params.Age, params.ProfilePictureUrl, params.ID)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "user already exists"})
		}
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "User updated successfully"})
}
