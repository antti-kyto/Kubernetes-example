package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if os.Getenv("EMOJI") == "" {
		os.Setenv("EMOJI", "🌍")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello")
		return c.SendString(fmt.Sprintf("Hello World %s!", os.Getenv("EMOJI")))
	})

	app.Get("/stop", func(c *fiber.Ctx) error {
		fmt.Println("Stopping the Server")
		os.Exit(0)
		return nil
	})

	app.Listen(":3000")
}
