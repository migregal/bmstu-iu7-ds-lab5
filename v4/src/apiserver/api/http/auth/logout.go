package auth

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

func (a *api) Logout(c echo.Context) error {
	logoutURL, err := url.Parse("https://" + a.authcfg.Domain + "/v2/logout")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	scheme := "http"
	if c.Request().TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + c.Request().Host)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", a.authcfg.ClientID)
	logoutURL.RawQuery = parameters.Encode()

	return c.Redirect(http.StatusTemporaryRedirect, logoutURL.String())
}
