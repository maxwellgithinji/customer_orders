package auth

import (
	"context"
	"log"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Authenticator is a struct that helps in OIDC auth
type Authenticator struct {
	Provider   *oidc.Provider
	Config     oauth2.Config
	Ctx        context.Context
	Verifier   *oidc.IDTokenVerifier
	OidcConfig *oidc.Config
}

// NewAuthenticator helps perform the authentication, amd returns error if it was unsuccessful
func NewAuthenticator() (*Authenticator, error) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")
	openIDDomainURL := os.Getenv("OPEN_ID_DOMAIN_URL")

	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, openIDDomainURL)
	if err != nil {
		log.Printf("Failed to get provider: %v\n", err)
		return nil, err
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)
	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	return &Authenticator{
		Provider:   provider,
		Config:     conf,
		Ctx:        ctx,
		Verifier:   verifier,
		OidcConfig: oidcConfig,
	}, nil
}
