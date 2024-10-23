package encryption

import (
	"picadosYa/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func SignedLoginToken(u *models.User) (string, error) {
	// Es viable este método si el servidor que creó el token
	// es el que se encarga de validarlo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      u.Email,
		"first_name": u.FirstName,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
		"role":       u.Role,
	})

	// firma el jwt
	return token.SignedString([]byte(key))

}
