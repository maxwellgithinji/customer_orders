package authcontroller

import (
	"net/http"
	"net/url"
	"os"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	clientID := os.Getenv("CLIENT_ID")
	openIDDomain := os.Getenv("OPEN_ID_DOMAIN")

	logoutUrl, err := url.Parse("https://" + openIDDomain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", clientID)
	logoutUrl.RawQuery = parameters.Encode()

	// TODO: ensure the session is invalidated after logout.
	// TODO: if a user tries to access a resource with an invalidated route, they are redirected to homepage
	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}
