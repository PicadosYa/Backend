package middlewares

import (
	"net/http"
	"picadosYa/internal/service"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	Exp  int64  `json:"exp"`
	jwt.StandardClaims
}

func JWTMiddleware(secretKey []byte, userService service.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract token from Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Missing authorization token",
				})
			}

			// Remove "Bearer " prefix
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse and validate token
			token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token",
				})
			}

			// Check token validity
			claims, ok := token.Claims.(*CustomClaims)
			if !ok || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token claims",
				})
			}

			// Check token expiration
			if claims.Exp < time.Now().Unix() {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Token expired",
				})
			}

			// Check if user exists in database
			_, err = userService.GetUserByID(c.Request().Context(), claims.ID)
			if err != nil {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": "User not found",
				})
			}
			// Store claims in context for route handlers
			c.Set("user", claims)

			return next(c)
		}
	}
}

// Role-based access control middleware
func RequireRole(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, ok := c.Get("user").(*CustomClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "No authentication context",
				})
			}

			// Check if user's role is in allowed roles
			for _, role := range allowedRoles {
				if claims.Role == role {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "Insufficient permissions",
			})
		}
	}
}
