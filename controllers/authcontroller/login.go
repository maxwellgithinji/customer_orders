package authcontroller

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/auth"
)

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

	// TODO: If a customer does not exist, save them in the db, otherwise skip saving
	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}
