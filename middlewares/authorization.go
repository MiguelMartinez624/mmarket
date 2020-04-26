package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/miguelmartinez624/mmarket/modules/authentication/core"
	sm "github.com/miguelmartinez624/mmarket/modules/stores/core"
)

var authModule *auth.Module
var storesModule *sm.Module

func SetAuthModule(module *auth.Module) {
	authModule = module
}
func SetStoresModule(module *sm.Module) {
	storesModule = module
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
func OwnStore(callback func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := r.Context().Value("user"); user != nil {

			profileID := user.(*auth.TokenClaims).ProfileID
			storeID := mux.Vars(r)["store_id"]

			_, err := storesModule.GetStoreByIDAndProfileID(r.Context(), storeID, profileID)
			if err != nil {

				fmt.Fprintf(w, "no authorise to resource")

				return
			}

			callback(w, r)

		}
	})
}
