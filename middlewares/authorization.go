package middlewares

import (
	"context"
	"fmt"
	"net/http"

	auth "github.com/miguelmartinez624/mmarket/modules/authentication/core"
)

var authModule *auth.Module

func SetAuthModule(module *auth.Module) {
	authModule = module
}

func IsAuthorozed(callback func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token, ok := r.Header["Token"]; ok {
			// Parse the token
			id, err := authModule.ValidateToken(r.Context(), token[0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			//Add data to context
			ctx := context.WithValue(r.Context(), "accountID", id)
			callback(w, r.WithContext(ctx))
		} else {
			fmt.Fprintf(w, "no aturoize")
		}
	})

}
