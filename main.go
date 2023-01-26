package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ssrahul96/go-echo/models"
)

func main() {
	app := fiber.New()

	app.All("/", func(c *fiber.Ctx) error {
		now := time.Now()
		var data models.LogModel

		data.Timestamp = now.UnixNano()
		data.Headers = c.GetReqHeaders()
		data.Url = c.Path()
		data.Method = c.Method()
		data.Body = string(c.Body())

		addCont := os.Getenv("ADDITIONAL_CONTENT")
		if len(addCont) != 0 {
			data.AdditionalContents = addCont
		}

		logData, _ := json.Marshal(data)
		log.Print(string(logData))
		return c.JSON(data)
	})

	app.Listen(":80")
}
