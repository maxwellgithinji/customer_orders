package middlewares

import (
	"context"
	"net/http"

	"github.com/maxwellgithinji/customer_orders/auth"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := auth.Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, ok := session.Values["profile"]; !ok {
			http.Redirect(w, r, "/api/v1", http.StatusSeeOther)
		} else {
			// Enable XSS protection with http only
			session.Options.HttpOnly = true
			session.Options.Secure = r.TLS != nil
			err = session.Save(r, w)
			next.ServeHTTP(w, r.WithContext(context.Background()))
		}
	})
}
