package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"os"
)

var photos []os.DirEntry

func main() {
	var err error
	photos, err = os.ReadDir("images")
	if err != nil {
		fmt.Println("Could not read cat photo folder!")
	}
	server()
}

func getRandomCatPic() string {
	return "images/" + photos[rand.Intn(len(photos))].Name()
}

func server() {
	app := fiber.New()

	app.Get("/cat", func(c *fiber.Ctx) error {
		return c.SendFile("home.html")
	})

	app.Get("/cat/img.png", func(c *fiber.Ctx) error {
		return c.SendFile(getRandomCatPic())
	})

	app.Listen(":8080")
}
