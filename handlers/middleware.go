package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func JWTAuthen() fiber.Handler {
	return func(c *fiber.Ctx) error {
		hmacSampleSecret := []byte(os.Getenv(viper.GetString("JWT_SECRET")))
		header := c.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})

		if err != nil {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "forbidden", "message": err.Error()})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("uuid", claims["uuid"])
		} else {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "forbidden", "message": "Invalid token"})
		}

		return c.Next()
	}
}
