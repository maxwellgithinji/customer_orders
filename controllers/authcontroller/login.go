package authcontroller

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/auth"
)

type login = struct {
	Redirect string
}

// Login redirects a user to authorize with OpenID connect
// @Summary Gets the redirect url for OpenID Login
// @Description redirects a user to authorize with OpenID connect
// @Tags  Auth
// @Produce  json
// @Success 200 {object} login{}
// @Router /login [get]
func Login(w http.ResponseWriter, r *http.Request) {
	// Generate random state
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := auth.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result = login{}
	result.Redirect = authenticator.Config.AuthCodeURL(state)
	json.NewEncoder(w).Encode(result)

	// TODO: If a customer does not exist, save them in the db, otherwise skip saving
	// http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}
