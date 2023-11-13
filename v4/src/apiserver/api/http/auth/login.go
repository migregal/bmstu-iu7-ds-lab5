package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (a *api) Login(c echo.Context) error {
	state, err := generateRandomState()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	sess, err := session.Get(sessionCookie, c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	sess.Options = &sessions.Options{
		Path:     "/oauth2/v1/auth0",
		MaxAge:   int(5 + time.Minute.Seconds()), //nolint: gomnd
		HttpOnly: true,
	}

	sess.Values["state"] = state
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, a.auth.AuthCodeURL(state))
}

func generateRandomState() (string, error) {
	b := make([]byte, 32) //nolint: gomnd
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate state: %w", err)
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
