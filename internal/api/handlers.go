package api

import (
	"log"
	"net/http"
	"picadosYa/encryption"
	"picadosYa/internal/api/dtos"
	"picadosYa/internal/models"
	"picadosYa/internal/service"
	"picadosYa/utils"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

func (a *API) GetFavouritesPerUser(c echo.Context) error {
	ctx := c.Request().Context()
	tokenStr := c.Request().Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	claims, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}
	id_user, ok1 := claims["id"].(float64)
	if ok1 != true {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Check id_user"})
	}
	idUser := int(id_user)
	favouritesPerUser, err := a.serv.GetFavouritesPerUser(ctx, idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, favouritesPerUser)
}

func (a *API) CreateOrRemoveFavourite(c echo.Context) error {
	ctx := c.Request().Context()
	tokenStr := c.Request().Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	favStruct := models.Fav{}
	err := c.Bind(&favStruct)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	claims, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}
	id_user, ok1 := claims["id"].(float64)
	if ok1 != true {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Check id_user"})
	}
	idUser := int(id_user)
	a.serv.CreateOrRemoveFavourite(ctx, idUser, favStruct.FieldID)
	return c.NoContent(http.StatusOK)
}

func (a *API) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	tokenStr := c.Request().Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	claims, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}
	id_user, ok1 := claims["id"].(float64)
	if ok1 != true {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Check id_user"})
	}
	idUser := int(id_user)
	user, err := a.serv.GetUserByID(ctx, idUser)

	userToFront := models.User{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		Phone:             user.Phone,
		ProfilePictureUrl: user.ProfilePictureUrl,
		Role:              user.Role,
		PositionPlayer:    user.ProfilePictureUrl,
		Age:               user.Age,
		IsVerified:        user.IsVerified,
	}
	return c.JSON(http.StatusOK, userToFront)

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
	form, err := c.MultipartForm()
	log.Printf("form: %v", form)
	// Parsear los datos del formulario
	params := new(dtos.UpdateUser)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	// Validar los parámetros del DTO
	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	// Obtener el archivo de imagen de perfil
	file, err := c.FormFile("profilePicture")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Error processing profile picture"})
	}

	var profilePictureURL string
	if file != nil {
		// Si se proporcionó un archivo, subirlo
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error opening file"})
		}
		defer src.Close()

		// Llamar al servicio con el archivo
		profilePictureURL, err = a.serv.UpdateUserInfo(ctx, params.FirstName, params.LastName, params.Email,
			params.Phone, params.PositionPlayer, params.TeamName, params.Age, file, params.ID)
		if err != nil {
			if err == service.ErrUserAlreadyExists {
				return c.JSON(http.StatusConflict, responseMessage{Message: "user already exists"})
			}
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
		}
	} else {
		// Si no se proporcionó archivo, llamar al servicio sin archivo
		profilePictureURL, err = a.serv.UpdateUserInfo(ctx, params.FirstName, params.LastName, params.Email,
			params.Phone, params.PositionPlayer, params.TeamName, params.Age, nil, params.ID)
		if err != nil {
			if err == service.ErrUserAlreadyExists {
				return c.JSON(http.StatusConflict, responseMessage{Message: "user already exists"})
			}
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
		}
	}

	return c.JSON(http.StatusOK, dtos.UpdateUser{
		ID:                params.ID,
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		Phone:             params.Phone,
		PositionPlayer:    params.PositionPlayer,
		TeamName:          params.TeamName,
		Age:               params.Age,
		ProfilePictureUrl: profilePictureURL,
	},
	)
}
