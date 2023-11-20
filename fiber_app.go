package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kataras/golog"
)

const profileUnknown = "unknown"

func main() {
	webApp := fiber.New()
	webApp.Get("/Users/dbelyashov/golang/myProject/profiles", func(c *fiber.Ctx) error {
		profileID := c.Query("profile_id", profileUnknown)
		if profileID == "" {
			profileID = profileUnknown
		}

		if profileID == profileUnknown {
			return c.Status(http.StatusUnprocessableEntity).SendString("profile_id is required")
		}

		return c.SendString(fmt.Sprintf("User Profile ID: %s", profileID))
	})

	golog.Fatal(webApp.Listen(":80"))
}
