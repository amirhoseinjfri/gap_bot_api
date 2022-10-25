package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return nil
	})

	app.Listen(":3000")
}
