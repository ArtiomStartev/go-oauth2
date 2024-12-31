# OAuth 2.0 Implementation in Golang

This repository contains an implementation of OAuth 2.0 using Golang, leveraging the Fiber web framework and the `golang.org/x/oauth2` package. The project demonstrates authentication and authorization using Google and GitHub APIs.

## Introduction
Security is a critical aspect of any service or API. This project provides a practical demonstration of implementing OAuth 2.0, an open authorization standard designed to grant applications limited access to user resources on behalf of the user.

## Key Features
- OAuth 2.0 implementation using Google and GitHub as authentication providers.
- Use of `golang.org/x/oauth2` for managing OAuth 2.0 flow.
- Fiber framework for building the API.

## Prerequisites
- Golang installed ([Download and Install](https://golang.org/dl/))
- Fiber web framework installed: `go get -u github.com/gofiber/fiber/v2`
- Client credentials from Google and GitHub APIs.

## Setup

### 1. Clone the Repository
```bash
git clone https://github.com/ArtiomStartev/go-oauth2.git
cd go-oauth2
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Create the `.env` File
Create a `.env` file in the root directory and populate it with your Google and GitHub API credentials:

```env
GOOGLE_CLIENT_ID=<your_google_client_id>
GOOGLE_CLIENT_SECRET=<your_google_client_secret>
GOOGLE_STATE=randomstate

GITHUB_CLIENT_ID=<your_github_client_id>
GITHUB_CLIENT_SECRET=<your_github_client_secret>
GITHUB_STATE=randomstate
```

### 4. Run the Application
```bash
go run main.go
```

## Endpoints

### Google OAuth 2.0
- **GET `/google_login`**: Redirects the user to Google's login page.
- **GET `/google_callback`**: Handles the callback from Google after successful login and retrieves user information.

### GitHub OAuth 2.0
- **GET `/github_login`**: Redirects the user to GitHub's login page.
- **GET `/github_callback`**: Handles the callback from GitHub after successful login and retrieves user information.

## Project Structure
```plaintext
├── config
│   ├── config.go          # OAuth configurations for Google and GitHub
├── controller
│   ├── google.go          # Handlers for Google OAuth 2.0
│   ├── github.go          # Handlers for GitHub OAuth 2.0
├── routes
│   ├── routes.go          # API route setup
├── .env                   # Environment variables
├── main.go                # Entry point of the application
```

## How It Works
1. **GoogleConfig and GitHubConfig**:
    - Reads client credentials and state from the `.env` file.
    - Sets up `oauth2.Config` for each provider with the required scopes and endpoints.

2. **OAuth 2.0 Flow**:
    - The user initiates login by visiting `/google_login` or `/github_login`.
    - The application redirects the user to the provider's login page.
    - Upon successful login, the provider redirects the user back to `/google_callback` or `/github_callback` with an authorization code.
    - The application exchanges the authorization code for an access token and retrieves user information.

## Example Usage
Visit the following URL in your browser to test the Google OAuth 2.0 flow:
```
http://127.0.0.1:8080/google_login
```
Upon successful login, user information will be displayed.

## Dependencies
- [Fiber](https://github.com/gofiber/fiber): Web framework for Golang.
- [OAuth2](https://pkg.go.dev/golang.org/x/oauth2): OAuth 2.0 library for Go.
- [godotenv](https://github.com/joho/godotenv): Load environment variables from a `.env` file.
