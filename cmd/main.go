package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {

	app := fiber.New()

	app.Get("/healthcheck", healthcheck)

	fmt.Println("Hello, World!")

	log.Fatal(app.Listen(":3000"))
}
