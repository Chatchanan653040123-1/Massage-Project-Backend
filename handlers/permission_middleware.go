package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h userHandler) AdminPermissionLevel1() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uuid, err := uuid.Parse(c.Locals("uuid").(string))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "forbidden", "message": "failed to parse uuid"})
		}
		permission, err := h.userSrv.GetEntityPermission(uuid)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "forbidden", "message": "failed to check admin permission"})
		}
		fmt.Println(permission)
		if permission.Role == "admin" && permission.PermissionLevel >= 1 {
			return c.Next()
		} else {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "forbidden", "message": "You are not normal admin"})
		}
	}
}
func (h userHandler) AdminPermissionLevel2() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uuid, err := uuid.Parse(c.Locals("uuid").(string))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "forbidden", "message": "failed to parse uuid"})
		}
		permission, err := h.userSrv.GetEntityPermission(uuid)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "forbidden", "message": "failed to check admin permission"})
		}
		fmt.Println("\n\n", permission)
		if permission.Role == "admin" && permission.PermissionLevel >= 2 {
			return c.Next()
		} else {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "forbidden", "message": "You are not super admin"})
		}
	}
}
func (h userHandler) UserPermissionLevel1() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uuid, err := uuid.Parse(c.Locals("uuid").(string))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "forbidden", "message": "failed to parse uuid"})
		}
		permission, err := h.userSrv.GetEntityPermission(uuid)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "forbidden", "message": "failed to check admin permission"})
		}
		fmt.Println(permission)
		if permission.Role == "admin" || (permission.Role == "user" && permission.PermissionLevel >= 1) {
			return c.Next()
		} else {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "forbidden", "message": "Your account got banned"})
		}
	}
}
