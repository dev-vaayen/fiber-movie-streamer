package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/films/:name", StreamFilm)

	app.Listen(":8000")
}

func StreamFilm(c *fiber.Ctx) error {
	filmDir := "./films"
	f := c.Params("name")

	if f == "" {
		return c.Status(400).JSON(fiber.Map{"error": "filename field required in params"})
	}

	fp := filepath.Join(filmDir, f+".mp4")
	fmt.Println(fp)

	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "Movie does not exist"})
	}

	c.Set("Content-Type", "video/mp4")
	return c.SendFile(fp)
}
