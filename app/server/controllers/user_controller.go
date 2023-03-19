package controllers

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinloaiza12/roller-tempo/app/database"
)

func Users(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": "hola front, saludos desde el back",
	})
}

func UserInfo(ctx context.Context, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("ID no válido")
		}
		result, err := database.GetUserByID(ctx, db, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener información del usuario")
		}
		return c.JSON(result.UserToJSON())
	}
}