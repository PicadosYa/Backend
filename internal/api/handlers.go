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

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid request"})
	}

	// valida lo que tenemos asignado en el dto
	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.FirstName, params.Lastname, params.Email, params.Password, params.Phone, params.Role, params.AcceptedTerms)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, models.ResponseMessage{Message: "user already exists"})
		}
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "internal server error"})
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
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid request"})
	}
	log.Println(err)
	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}
	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: err.Error()})
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
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  userCreated,
		"token": token,
	})
}

func (a *API) GetFavouritesPerUser(c echo.Context) error {
	ctx := c.Request().Context()
	idUser := utils.GenerateUserID(c)
	favouritesPerUser, err := a.serv.GetFavouritesPerUser(ctx, idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, favouritesPerUser)
}

func (a *API) CreateOrRemoveFavourite(c echo.Context) error {
	ctx := c.Request().Context()
	favStruct := models.Fav{}
	err := c.Bind(&favStruct)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid request"})
	}
	idUser := utils.GenerateUserID(c)
	a.serv.CreateOrRemoveFavourite(ctx, idUser, favStruct.FieldID)
	return c.NoContent(http.StatusOK)
}

func (a *API) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	idUser := utils.GenerateUserID(c)
	user, err := a.serv.GetUserByID(ctx, idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: err.Error()})
	}
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
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid request"})
	}

	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	err = a.serv.ResetPassword(ctx, params.Email, params.Token, params.NewPassword)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, models.ResponseMessage{Message: "user already exists"})
		}
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, models.ResponseMessage{Message: "Password successfully updated"})
}

func (a *API) GetExpiration(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	tkn, err := encryption.ParseLoginJWT(tokenStr)
	if err != nil {
		log.Println(tokenStr)
		log.Println(tkn)
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	expiresVal, ok := tkn["exp"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "Formato de token inválido"})
	}

	expirationUnix := int64(expiresVal)
	expirationTime := time.Unix(expirationUnix, 0)

	timeRemaining := time.Until(expirationTime)
	if timeRemaining <= 0 {
		return c.JSON(http.StatusUnauthorized, models.ResponseMessage{Message: "El token ha expirado"})
	}

	return c.JSON(http.StatusOK, models.ResponseMessage{Message: "Ok"})
}

func (a *API) VerifyUserEmail(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RequestSendEmail{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid request"})
	}

	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	// Genera el token de 6 digitos
	recoveryToken := utils.GenerateRandomDigits(6)

	err := a.serv.SaveToken(ctx, params.Email, recoveryToken, time.Now().Add(15*time.Minute))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "Unable to save recovery token"})
	}

	// Envía el mail
	err = a.serv.SendVerifyEmail(params.Email, recoveryToken)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.ResponseMessage{Message: "Verify email sent"})
}

func (a *API) UpdateVerifyUser(c echo.Context) error {
	ctx := c.Request().Context()
	token := c.QueryParam("token")

	// Verificar el token en la base de datos
	user, err := a.serv.GetUserByToken(ctx, token)
	if err != nil || user == nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	// Actualizar isVerified a true
	err = a.serv.UpdateUserVerification(ctx, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.ResponseMessage{Message: "User updated successfully"})
}

func (a *API) RequestPasswordRecovery(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RequestSendEmail{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid request"})
	}

	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	// Genera el token de 6 digitos
	recoveryToken := utils.GenerateRandomDigits(6)

	err := a.serv.SaveToken(ctx, params.Email, recoveryToken, time.Now().Add(15*time.Minute))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "Unable to save recovery token"})
	}

	// Envía el mail
	err = a.serv.SendRecoveryEmail(params.Email, recoveryToken)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.ResponseMessage{Message: "Recovery email sent"})
}

func (a *API) UpdateUserProfileInfo(c echo.Context) error {
	ctx := c.Request().Context()
	form, err := c.MultipartForm()
	log.Printf("form: %v", form)
	// Parsear los datos del formulario
	params := new(dtos.UpdateUser)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Invalid request"})
	}

	// Validar los parámetros del DTO
	if err := a.dataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: err.Error()})
	}

	// Obtener el archivo de imagen de perfil
	file, err := c.FormFile("profilePicture")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, models.ResponseMessage{Message: "Error processing profile picture"})
	}

	//var profilePictureURL string
	if file != nil {
		// Si se proporcionó un archivo, subirlo
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "Error opening file"})
		}
		defer src.Close()

		// Llamar al servicio con el archivo
		_, err = a.serv.UpdateUserInfo(ctx, params.FirstName, params.LastName, params.Email,
			params.Phone, params.PositionPlayer, params.TeamName, params.Age, file, params.ID, "")
		if err != nil {
			if err == service.ErrUserAlreadyExists {
				return c.JSON(http.StatusConflict, models.ResponseMessage{Message: "user already exists"})
			}
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "internal server error"})
		}
	} else {
		// Si no se proporcionó archivo, llamar al servicio sin archivo
		_, err = a.serv.UpdateUserInfo(ctx, params.FirstName, params.LastName, params.Email,
			params.Phone, params.PositionPlayer, params.TeamName, params.Age, nil, params.ID, params.ProfilePictureUrl)
		if err != nil {
			if err == service.ErrUserAlreadyExists {
				return c.JSON(http.StatusConflict, models.ResponseMessage{Message: "user already exists"})
			}
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "internal server error"})
		}
	}

	user, err := a.serv.GetUserByID(ctx, int(params.ID))

	token, err := encryption.SignedLoginToken(&models.User{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		Phone:             user.Phone,
		ProfilePictureUrl: user.ProfilePictureUrl,
		Role:              user.Role,
		PositionPlayer:    user.PositionPlayer,
		Age:               user.Age,
		IsVerified:        user.IsVerified,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "token error"})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (a *API) RefreshToken(c echo.Context) error {
    ctx := c.Request().Context()
    
    // Obtener el ID del usuario desde el token
    idUser := utils.GenerateUserID(c)

    // Obtener la información completa del usuario
    user, err := a.serv.GetUserByID(ctx, idUser)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "Error retrieving user"})
    }

    // Generar un nuevo token
    token, err := encryption.SignedLoginToken(&models.User{
        ID:                user.ID,
        FirstName:         user.FirstName,
        LastName:          user.LastName,
        Email:             user.Email,
        Phone:             user.Phone,
        ProfilePictureUrl: user.ProfilePictureUrl,
        Role:              user.Role,
        PositionPlayer:    user.PositionPlayer,
        Age:               user.Age,
        IsVerified:        user.IsVerified,
    })
    if err != nil {
        return c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: "Error generating token"})
    }

    // Devolver el nuevo token
    return c.JSON(http.StatusOK, map[string]string{
        "token": token,
    })
}
