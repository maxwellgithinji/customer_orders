package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/customer_orders/services/openidauthservice"
	"github.com/maxwellgithinji/customer_orders/utils"
)

type AuthMiddleware interface {
	IsAuthenticated(next http.Handler) http.Handler
	IsCurrentUser(next http.Handler) http.Handler
}

type authmiddleware struct{}

var (
	openIDAuthService openidauthservice.OpenIdAuthService
)

func NewAuthMiddleware(openIdAuth openidauthservice.OpenIdAuthService) AuthMiddleware {
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
			utils.ResponseHelper(w, "500", err.Error()+",  delete the session cookie")
			return
		}
		if _, ok := session.Values["profile"]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseHelper(w, "401", "Unauthorized. Please log in")
			return
		} else {
			// Enable XSS protection with http only
			session.Options.HttpOnly = true
			session.Options.Secure = r.TLS != nil
			_ = session.Save(r, w)
			ctx := context.WithValue(r.Context(), "auth-session", session)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func (*authmiddleware) IsCurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		err := openIDAuthService.InitSession()
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
		sessionemail := fmt.Sprintf("%v", session.Values["profile"].(map[string]interface{})["email"])
		if _, ok := session.Values["profile"]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseHelper(w, "401", "Unauthorized. Please log in")
			return
		} else if sessionemail != params["email"] {
			w.WriteHeader(http.StatusUnauthorized)
			utils.ResponseHelper(w, "401", "Unauthorized. Only current user allowed")
			return
		} else {
			// Enable XSS protection with http only
			session.Options.HttpOnly = true
			session.Options.Secure = r.TLS != nil
			_ = session.Save(r, w)
			ctx := context.WithValue(r.Context(), "auth-session", session)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
