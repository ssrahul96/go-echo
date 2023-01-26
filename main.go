package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ssrahul96/go-echo/models"
	"github.com/ssrahul96/golang/utils"
)

func main() {
	app := fiber.New()

	app.All("/", func(c *fiber.Ctx) error {
		now := time.Now()
		var data models.LogModel
		data.Timestamp = now.Unix()
		data.Headers = c.GetReqHeaders()
		data.Url = c.BaseURL()
		data.Method = c.Method()
		return c.SendString(utils.FormatJson(data))
	})

	app.Listen(":3000")
}
