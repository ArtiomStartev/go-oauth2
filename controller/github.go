package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtiomStartev/go-oauth2/config"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
)

func GitHubLogin(c *fiber.Ctx) error {
	url := config.AppConfig.GitHubLoginConfig.AuthCodeURL(os.Getenv("GITHUB_STATE"))

	c.Status(fiber.StatusSeeOther)
	if err := c.Redirect(url); err != nil {
		fmt.Println("Error redirecting to GitHub login: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error redirecting to GitHub login",
		})
	}

	return c.JSON(url)
}

func GitHubCallback(c *fiber.Ctx) error {
	if state := c.Query("state"); state != os.Getenv("GITHUB_STATE") {
		fmt.Println("Invalid state: ", state)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid state",
		})
	}

	authCode := c.Query("code")
	token, err := config.AppConfig.GitHubLoginConfig.Exchange(context.Background(), authCode)
	if err != nil {
		fmt.Println("Error exchanging auth code for token: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error exchanging auth code for token",
		})
	}

	res, err := http.Get("https://api.github.com/user/repo?access_token=" + token.AccessToken)
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Println("Error getting user info: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting user info",
		})
	}
	defer res.Body.Close()

	var userData map[string]any
	if err = json.NewDecoder(res.Body).Decode(&userData); err != nil {
		fmt.Println("Error decoding user data: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error decoding user data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(userData)
}
