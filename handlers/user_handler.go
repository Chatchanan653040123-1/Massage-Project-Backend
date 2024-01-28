package handlers

import (
	"massage/logs"
	"massage/services"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	registerLog := "User " + response.UUID.String() + " has been registered"
	logs.Info(registerLog)
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
		loginLog := "User " + response.UUID.String() + " has been logged in"
		logs.Info(loginLog)
		return c.JSON(fiber.Map{"token": tokenString})

	} else {
		loginLog := "User " + response.UUID.String() + " can't log in because password is incorrect\ncompare: " + compare.Error() + "\nIncorrect password: " + request.Password
		logs.Error(loginLog)
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
func (h userHandler) GetMyAccount(c *fiber.Ctx) error {
	uuid, err := uuid.Parse(c.Locals("uuid").(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to parse uuid",
		})
	}
	user, err := h.userSrv.GetUser(uuid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to get user",
		})
	}

	return c.JSON(user)
}
func (h userHandler) UpdateMyAccount(c *fiber.Ctx) error {
	uuid, err := uuid.Parse(c.Locals("uuid").(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to parse uuid",
		})
	}
	request := services.RegisterBody{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}
	user, err := h.userSrv.UpdateMyAccount(uuid, request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to update user",
		})
	}

	return c.JSON(user)
}
