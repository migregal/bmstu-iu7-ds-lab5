package authenticator

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Middleware struct {
	lg *slog.Logger

	auth *Authenticator
}

func NewMiddleware(lg *slog.Logger, auth *Authenticator) *Middleware {
	return &Middleware{lg: lg, auth: auth}
}

func (m *Middleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")

		token, found := strings.CutPrefix(auth, "Bearer ")
		if !found {
			m.lg.Error("failed to obtain auth header")

			return c.NoContent(http.StatusUnauthorized)
		}

		idToken, err := m.auth.verifyIDToken(c.Request().Context(), token)
		if err != nil {
			m.lg.Error("failed to verify id token", "err", err)

			return c.NoContent(http.StatusUnauthorized)
		}

		tmp := struct {
			Email string `json:"email"`
		}{}
		if err = idToken.Claims(&tmp); err != nil {
			m.lg.Error("failed to parse claims", "err", err)

			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Set("username", tmp.Email)

		return next(c)
	}
}
