package service

import (
	"context"
	"encoding/gob"
	"log"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

type OpenIdAuthService interface {
	Authenticate(openIDDomainURL string) (*Authenticator, error)
	InitSession() error
	NewStore() *sessions.FilesystemStore
}

// Authenticator is a struct that helps in OIDC auth
type Authenticator struct {
	Provider   *oidc.Provider
	Config     oauth2.Config
	Ctx        context.Context
	Verifier   *oidc.IDTokenVerifier
	OidcConfig *oidc.Config
}

type auth struct {
}

var (
	Store   *sessions.FilesystemStore
	Options *sessions.Options
)

// NewOpenIdAuthService
func NewOpenIdAuthService() OpenIdAuthService {
	return &auth{}
}

func (*auth) Authenticate(openIDDomainURL string) (*Authenticator, error) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")
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

func (*auth) InitSession() error {
	sessionKey := os.Getenv("SESSION_KEY")
	Store = sessions.NewFilesystemStore("", []byte(sessionKey))
	Options = &sessions.Options{
		MaxAge:   int(time.Hour.Seconds()),
		HttpOnly: true,
	}
	gob.Register(map[string]interface{}{})
	return nil
}

func (*auth) NewStore() *sessions.FilesystemStore {
	return Store
}
