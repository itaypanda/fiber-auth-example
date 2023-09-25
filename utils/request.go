package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/itaypanda/fiber-auth-example/models"
	"github.com/itaypanda/fiber-auth-example/database"
	"github.com/gofiber/fiber/v2"
)

func GetUserFromContext(c *fiber.Ctx) (*models.User, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	dbUser, err := database.FindUser(id)
	if err != nil {
		return dbUser, err
	}
	return dbUser, nil
}
