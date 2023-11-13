package auth

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type CallbackResponse struct {
	AccessToken string `json:"access_token"`
}

func (a *api) Callback(c echo.Context) error {
	sess, err := session.Get(sessionCookie, c)
	if err != nil {
		return c.String(http.StatusBadRequest, "no session")
	}

	if state, ok := sess.Values["state"].(string); !ok || c.QueryParam("state") != state {
		return c.String(http.StatusBadRequest, "invalid state")
	}

	token, err := a.auth.Exchange(c.Request().Context(), c.QueryParam("code"))
	if err != nil {
		return c.String(http.StatusUnauthorized, "failed to convert an authorization code into a token.")
	}

	_, err = a.auth.VerifyIDTokenFromTokenSet(c.Request().Context(), token)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to verify ID Token.")
	}

	resp := CallbackResponse{AccessToken: token.Extra("id_token").(string)} //nolint: forcetypeassert

	return c.JSON(http.StatusOK, resp)
}
