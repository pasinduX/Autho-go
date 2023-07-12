package login

import (
	"AUTH0_FIBER/platform/authenticator"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
)

// Handler for our login.
func Handler(auth *authenticator.Authenticator) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		state, err := generateRandomState()
		if err != nil {
			return ctx.SendString(err.Error())
		}

		// Save the state inside the session.
		session := session.New().Get(ctx)
		session.Set("state", state)
		if err := session.Save(); err != nil {
			return ctx.SendString(err.Error())
		}

		return ctx.Redirect(auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}