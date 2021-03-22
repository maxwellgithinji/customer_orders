package middlewares

import (
	"context"
	"net/http"

	service "github.com/maxwellgithinji/customer_orders/services"
	"github.com/maxwellgithinji/customer_orders/utils"
)

type AuthMiddleware interface {
	IsAuthenticated(next http.Handler) http.Handler
}

type authmiddleware struct{}

var (
	openIDAuthService service.OpenIdAuthService
)

func NewAuthMiddleware(openIdAuth service.OpenIdAuthService) AuthMiddleware {
	openIDAuthService = openIdAuth
	return &authmiddleware{}
}

func (*authmiddleware) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		if _, ok := session.Values["profile"]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseHelper(w, "401", "Unauthorized. Please log in")
			// http.Redirect(w, r, "/api/v1", http.StatusSeeOther)
		} else {
			// Enable XSS protection with http only
			session.Options.HttpOnly = true
			session.Options.Secure = r.TLS != nil
			err = session.Save(r, w)
			next.ServeHTTP(w, r.WithContext(context.Background()))
		}
	})
}
