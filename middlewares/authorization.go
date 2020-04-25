package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/miguelmartinez624/mmarket/modules/authentication/core"
)

var authModule *auth.Module

func SetAuthModule(module *auth.Module) {
	authModule = module
}

func IsAuthorized(callback func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token, ok := r.Header["Token"]; ok {
			// Parse the token
			claims, err := authModule.ValidateToken(r.Context(), token[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			//Add data to context
			ctx := context.WithValue(r.Context(), "user", claims)
			// ctx = context.WithValue(r.Context(), "profileID", claims.ProfileID)
			callback(w, r.WithContext(ctx))
		} else {
			fmt.Fprintf(w, "no aturoize")
		}
	})
}

func OwnResource(callback func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := r.Context().Value("user"); user != nil {
			rProfile := mux.Vars(r)["profile_id"]
			tProfile := user.(*auth.TokenClaims).ProfileID
			//If not logged as the profile will be unhourtise
			// TODO make better error forthis.
			if rProfile != tProfile {

				fmt.Fprintf(w, "no authorise to resource")
				return
			}

			callback(w, r)

		}
	})
}
