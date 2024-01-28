package handlers

import (
	"massage/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h userHandler) GetUser(c *fiber.Ctx) error {
	uuid, err := uuid.Parse(c.Params("uuid"))
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
func (h userHandler) UpdateAccount(c *fiber.Ctx) error {
	uuid, err := uuid.Parse(c.Params("uuid"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to parse uuid",
		})
	}
	request := services.UpdateUserRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Body Parser",
		})
	}
	user, err := h.userSrv.UpdateAccount(uuid, request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to update user",
		})
	}

	return c.JSON(user)
}
func (h userHandler) DeleteAccount(c *fiber.Ctx) error {
	uuid, err := uuid.Parse(c.Params("uuid"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to parse uuid",
		})
	}
	_, err = h.userSrv.DeleteAccount(uuid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"Success": "User has been deleted",
	})
}
