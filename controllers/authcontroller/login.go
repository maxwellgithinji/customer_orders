package authcontroller

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/models"
	"github.com/maxwellgithinji/customer_orders/utils"
)

// Login redirects a user to authorize with OpenID connect
// @Summary Gets the redirect url for OpenID Login
// @Description redirects a user to authorize with OpenID connect
// @Tags  Auth
// @Produce  json
// @Success 200 {object} models.Login{}
// @Router /login [get]
func (*authcontroller) Login(w http.ResponseWriter, r *http.Request) {
	// Generate random state
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	state := base64.StdEncoding.EncodeToString(b)
	err = openIDAuthService.InitSession()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	session, err := openIDAuthService.NewStore().Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error()+",  delete the session cookie")
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
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
	var result = models.Login{}
	result.Redirect = authenticator.Config.AuthCodeURL(state)
	utils.ResponseWithDataHelper(w, "200", "Go to the redirect path below to log in", result)
	// http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}
