package encryption

import (
	"picadosYa/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func SignedLoginToken(u *models.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 5)
	// Es viable este método si el servidor que creó el token
	// es el que se encarga de validarlo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      u.Email,
		"first_name": u.FirstName,
		"exp":        expirationTime.Unix(),
		"role":       u.Role,
	})

	// firma el jwt
	return token.SignedString([]byte(key))
}

func ParseLoginJWT(value string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
