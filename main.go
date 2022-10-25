package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Text struct {
	ChatId int    `json:"chat_id"`
	Type   string `json:"type"`
	From   `json:"from"`
	Data   string `json:"data"`
}

type From struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	Username string `json:"user"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Body()))
		var text Text
		if err := json.Unmarshal(c.Body(), &text); err != nil {
			log.Println(err)
		}
		fmt.Println(text)
		return nil
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
