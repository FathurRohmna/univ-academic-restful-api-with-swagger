package pkgmiddleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token is required"})
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		studentID, ok := claims["student_id"].(string)
		if !ok || studentID == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token: student ID not found"})
		}

		c.Set("student_id", studentID)
		return next(c)
	}
}
