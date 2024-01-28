package handlers

import (
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
