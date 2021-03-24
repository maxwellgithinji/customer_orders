package authcontroller

import (
	"net/http"
	"net/url"
	"os"

	"github.com/maxwellgithinji/customer_orders/utils"
)

// LogOut redirects a user to log out with openID connect
// @Summary Gets the redirect url for openID connect logout
// @Description redirects a user to log out with openID connect
// @Tags  Auth
// @Produce  json
// @Success 200
// @Router /logout [post]
func (*authcontroller) Logout(w http.ResponseWriter, r *http.Request) {
	err := openIDAuthService.InitSession()
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
	clientID := os.Getenv("CLIENT_ID")
	openIDDomain := os.Getenv("OPEN_ID_DOMAIN")

	logoutUrl, err := url.Parse("https://" + openIDDomain)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	logoutUrl.Path += "/v2/logout"
	parameters := url.Values{}

	var scheme string
	if r.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host + "/api/v1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", clientID)
	logoutUrl.RawQuery = parameters.Encode()

	// Expire the session and save
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	_, err = http.Get(logoutUrl.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ResponseHelper(w, "500", err.Error())
		return
	}

	utils.ResponseHelper(w, "200", "Logout Successful")
	// http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)

}
