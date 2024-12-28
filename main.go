package main

import (
	"fmt"
	"github.com/ArtiomStartev/go-oauth2/config"
	"github.com/ArtiomStartev/go-oauth2/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if _, err := config.GoogleConfig(); err != nil {
		fmt.Println("Error setting up Google config: ", err)
		return
	}

	if _, err := config.GitHubConfig(); err != nil {
		fmt.Println("Error setting up GitHub config: ", err)
		return
	}

	routes.Setup(app)

	if err := app.Listen(":8080"); err != nil {
		fmt.Println("Error starting server on port 8080: ", err)
		return
	}
}
