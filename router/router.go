package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itaypanda/fiber-auth-example/handlers"
	"github.com/itaypanda/fiber-auth-example/utils"
)

func SetupPublicRotues(a *fiber.App) {
	api := a.Group("/api")

	auth := api.Group("/auth")

	auth.Post("/signup", handlers.SignupHandler)

	auth.Post("/signin", handlers.SigninHandler)
}

func SetupPrivateRotues(a *fiber.App) {
	api := a.Group("/api")

	api.Get("/aboutme", func(c *fiber.Ctx) error {
		user, err := utils.GetUserFromContext(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error fetching user",
			})
		}

		return c.JSON(fiber.Map{
			"name": user.Name,
			"email": user.Email,
		})
	})
}
