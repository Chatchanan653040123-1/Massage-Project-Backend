package handlers

import (
	"massage/services"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}
func (h userHandler) Registers(c *fiber.Ctx) error {
	request := services.RegisterBody{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}

	response, err := h.userSrv.Register(request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid Signup Credentials",
		})
	}

	return c.JSON(response)

}
func (h userHandler) Login(c *fiber.Ctx) error {
	request := services.LoginBody{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}

	if request.Identifier == "" || request.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Error": "invalid login credentials (Email or Password)",
		})
	}

	response, err := h.userSrv.Login(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "(Username or Email) and password is incorrect)",
		})
	}
	compare := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(request.Password))
	if compare == nil {
		hmacSampleSecret = []byte(os.Getenv(viper.GetString("JWT_SECRET")))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid": response.UUID,
			"iat":  time.Now().Unix(),
			"exp":  time.Now().Add(time.Hour * 1).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Error": "Failed to generate token"})
		}

		return c.JSON(fiber.Map{"token": tokenString})
	} else {
		return c.JSON(fiber.Map{"Error": "Password is incorrect"})
	}
}
func (h userHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userSrv.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to get all users",
		})
	}

	return c.JSON(users)
}
