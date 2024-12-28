package routes

import (
	"github.com/ArtiomStartev/go-oauth2/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/google_login", controllers.GoogleLogin)

	app.Get("/google_callback", controllers.GoogleCallback)

	app.Get("/github_login", controllers.GitHubLogin)

	app.Get("/github_callback", controllers.GitHubCallback)
}
