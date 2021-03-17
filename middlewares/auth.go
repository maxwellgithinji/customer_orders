package middlewares

import (
	"context"
	// "encoding/json"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/auth"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Initialize session
		err := auth.InitSession()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, err := auth.Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, ok := session.Values["profile"]; !ok {
			http.Redirect(w, r, "/api/v1", http.StatusSeeOther)
		} else {
			// profile := session.Values["profile"]
			// ctx := context.WithValue(r.Context(), "profile", profile)
			next.ServeHTTP(w, r.WithContext(context.Background()))
		}
	})
}
