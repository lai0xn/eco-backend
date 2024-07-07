package types

import "golang.org/x/oauth2"

type OAuthProvider struct {
	Config   *oauth2.Config
	Endpoint oauth2.Endpoint
}

var OAuth2Configs map[string]*OAuthProvider
