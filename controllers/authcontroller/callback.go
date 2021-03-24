package authcontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/models"
	"github.com/maxwellgithinji/customer_orders/utils"
)

func (*authcontroller) Callback(w http.ResponseWriter, r *http.Request) {
	err := openIDAuthService.InitSession()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	authenticator, err := openIDAuthService.Authenticate("https://maxgit.us.auth0.com/")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	session, err := openIDAuthService.NewStore().Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		w.WriteHeader(http.StatusBadRequest)
		utils.ResponseHelper(w, "400", "Invalid state parameter")
		return
	}

	token, err := authenticator.Config.Exchange(authenticator.Ctx, r.URL.Query().Get("code"))
	if err != nil {
		log.Printf("no token found: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		utils.ResponseHelper(w, "401", "no token found: "+err.Error())
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", "No id_token field in oauth2 token.")
		return
	}

	idToken, err := authenticator.Provider.Verifier(authenticator.OidcConfig).Verify(authenticator.Ctx, rawIDToken)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", "Failed to verify ID Token: "+err.Error())
		return
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	session.Options.HttpOnly = true
	session.Options.Secure = r.TLS != nil

	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	// userprofile := session.Values["profile"]
	email := fmt.Sprintf("%v", profile["email"])
	name := fmt.Sprintf("%v", profile["name"])
	status := "inactive"

	var customer models.Customer
	var defaultcustomerstate models.Customer
	var emailexists models.Customer

	customer.Username = name
	customer.Email = email
	customer.Status = status

	// Check if a user with current email exist before saving
	emailexists, err = customerService.FindACustomerByEmail(email)
	// TODO: debug pointers implementation to in order to check for nil instead of the code below
	if emailexists == defaultcustomerstate {
		if err == nil {
			// Save as new user
			_, err = customerService.CreateCustomer(customer)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				utils.ResponseHelper(w, "500", err.Error())
				return
			}
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	// Redirect to profile route after successful login
	defer http.Redirect(w, r, "/api/v1/auth/profile", http.StatusSeeOther)

}
