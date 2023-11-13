package auth

import (
	"fmt"
	"log/slog"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/migregal/bmstu-iu7-ds-lab2/pkg/oauth2/auth0/authenticator"
)

const sessionCookie = "library-auth"

type api struct {
	authcfg authenticator.Config
	auth    *authenticator.Authenticator
}

func InitListener(mx *echo.Echo, lg *slog.Logger, cfg authenticator.Config) error {
	gr := mx.Group("/oauth2/v1")

	gr.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	auth, err := authenticator.New(cfg)
	if err != nil {
		return fmt.Errorf("init auth: %w", err)
	}

	a := api{auth: auth}

	gr.GET("/auth0/authorize", a.Login)
	gr.GET("/auth0/callback", a.Callback)
	gr.GET("/auth0/logout", a.Logout)

	return nil
}
