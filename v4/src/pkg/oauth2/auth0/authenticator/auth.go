package authenticator

import (
	"context"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Config struct {
	Domain       string `mapstructure:"domain"`
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	CallbackURL  string `mapstructure:"callback_url"`
}

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// New instantiates the *Authenticator.
func New(cfg Config) (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+cfg.Domain+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

func (a *Authenticator) VerifyIDTokenFromTokenSet(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	return a.verifyIDToken(ctx, rawIDToken)
}

func (a *Authenticator) verifyIDToken(ctx context.Context, token string) (*oidc.IDToken, error) {
	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, token)
}
