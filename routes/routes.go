package routes

import (
	"github.com/ArtiomStartev/go-oauth2/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/google_login", controller.GoogleLogin)

	app.Get("/google_callback", controller.GoogleCallback)

	app.Get("/github_login", controller.GitHubLogin)

	app.Get("/github_callback", controller.GitHubCallback)
}
