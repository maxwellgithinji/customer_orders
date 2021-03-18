package authcontroller

import (
	// "context"
	"log"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/auth"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := auth.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	token, err := authenticator.Config.Exchange(authenticator.Ctx, r.URL.Query().Get("code"))
	if err != nil {
		log.Printf("no token found: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	idToken, err := authenticator.Provider.Verifier(authenticator.OidcConfig).Verify(authenticator.Ctx, rawIDToken)

	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	session.Options.HttpOnly = true
	session.Options.Secure = r.TLS != nil

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to profile route after successful login
	http.Redirect(w, r, "/api/v1/auth/profile", http.StatusSeeOther)
}
